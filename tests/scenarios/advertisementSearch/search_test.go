package advertisementSearch

import (
	"os"
	"strings"
	"testing"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Данные для проверки поиска объявления.
const (
	existingAdTitle = "оталвтм"
	searchBeginning = "от"
	searchMiddle    = "алв"
	searchEnd       = "втм"
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
	base.Precondition("Проверяем, что константы поиска заданы")

	require.NotEmpty(s.T(), existingAdTitle, "Задайте название объявления")
	require.NotEmpty(s.T(), searchBeginning, "Задайте начало названия")
	require.NotEmpty(s.T(), searchMiddle, "Задайте середину названия")
	require.NotEmpty(s.T(), searchEnd, "Задайте конец названия")

	s.loginData = auth.Login(
		s.T(),
		os.Getenv("TEST_LOGIN"),
		os.Getenv("TEST_PASSWORD"),
	)
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// SetupTest создаёт объявление для каждого теста поиска.
func (s *TestSuite) SetupTest() {
	t := s.T()
	t.Log("Подготовка: создаём объявление для теста поиска")

	advertisement := advList.CreateAdvertisement(
		t,
		s.loginData.Token,
		existingAdTitle,
		"Объявление для проверки поиска",
		"100",
		"1",
		"../../testdata/advertisement.jpg",
	)

	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления для поиска")

	t.Cleanup(func() {
		t.Log("Очистка: удаляем объявление для теста поиска")
		advList.DeleteAdvertisement(t, s.loginData.Token, advertisement.ID)
	})
}

func (s *TestSuite) TestSearchByTitle() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, searchMiddle)

	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})

	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
	})
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", searchBeginning},
		{"середина", searchMiddle},
		{"конец", searchEnd},
		{"нижний регистр", strings.ToLower(searchMiddle)},
		{"верхний регистр", strings.ToUpper(searchMiddle)},
		{"смешанный регистр", utils.MixCase(searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
		})
	}
}
