package myAdvertisement

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/myAdvertisements"

	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения и получение его параметров")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// TestGetMyAdvertisementsPositive проверяет получение созданного объявления.
func (s *TestSuite) TestGetMyAdvertisementsPositive() {
	t := s.T()
	t.Log("Подготовка: создаём собственное объявление")
	title := fmt.Sprintf("Мои объявления %d", time.Now().UnixNano())
	advertisement := advertisements.CreateAdvertisement(
		t, s.loginData.Token, title,
		"Объявление для проверки списка", "100", "1",
		"advertisement.jpg",
	)
	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления")
	t.Cleanup(func() {
		t.Log("Очистка: удаляем тестовое объявление")
		advertisements.DeleteAdvertisement(t, s.loginData.Token, advertisement.ID)
	})
	t.Log("Шаг: получаем список собственных объявлений")
	body := myAdvertisements.GetMyAdvertisements(t, s.loginData.Token, http.StatusOK, 50, title)
	items := gjson.Get(body, "items").Array()
	t.Log("Проверка: находим созданное объявление по ID")
	found := false
	for _, item := range items {
		if item.Get("id").String() != advertisement.ID {
			continue
		}
		found = true
		require.Equal(t, s.loginData.User.Id, item.Get("user_id").String(), "Владелец не совпадает")
		require.Equal(t, advertisement.Title, item.Get("title").String(), "Название не совпадает")
		require.NotEmpty(t, item.Get("photos").Array(), "У объявления нет фотографий")
		break
	}
	require.True(t, found, "Созданное объявление не найдено в списке собственных")
}

// Negative: Проверяем 401 ошибку при попытке получить собственные объявления без токена
func (s *TestSuite) TestGetMyAdvertisementsIncorrectToken() {
	var advertisementsBody string
	s.Run("Получаем список собственных объявлений с неправильным авторизационным токеном", func() {
		advertisementsBody = myAdvertisements.GetMyAdvertisements(s.T(), "incorrect_token", http.StatusUnauthorized, 50, "")
	})

	s.Run("Проверяем, что показывается правильная ошибка", func() {
		require.Equal(s.T(), "unauthorized", gjson.Get(advertisementsBody, "error").String(),
			"Некорректное значение error параметра")
		require.Equal(s.T(), "Invalid or expired token", gjson.Get(advertisementsBody, "message").String(),
			"Некорректное значение message параметра")
	})
}
