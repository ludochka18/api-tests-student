package advertisementCreate

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestCreateAdvertisementWithEmptyTitle(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуемся и готовим объявление без названия")
	login := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	t.Log("Шаг: отправляем запрос с пустым названием")
	apiError := advertisements.CreateAdvertisementExpectingError(
		t,
		login.Token,
		"",
		"Описание тестового объявления",
		"100",
		"1",
		testPhotoPath,
		http.StatusBadRequest,
	)

	t.Log("Проверка: сервер вернул сообщение об ошибке")
	require.Equal(t, "title is required", apiError.Message, "При пустом названии сервер вернул неожиданное сообщение")
}

func TestCreateAdvertisementWithEmptyDescription(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуемся и готовим объявление без описания")
	login := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	t.Log("Шаг: отправляем запрос с пустым описанием")
	apiError := advertisements.CreateAdvertisementExpectingError(
		t,
		login.Token,
		"Тестовое объявление",
		"",
		"100",
		"1",
		testPhotoPath,
		http.StatusBadRequest,
	)

	t.Log("Проверка: сервер вернул сообщение об ошибке")
	require.Equal(t, "description is required", apiError.Message, "При пустом описании сервер вернул неожиданное сообщение")
}

func TestCreateAdvertisementWithNegativePrice(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуемся и готовим объявление с отрицательной ценой")
	login := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	t.Log("Шаг: отправляем запрос с ценой -100")
	status, apiError := advertisements.CreateAdvertisementWithCleanup(
		t,
		login.Token,
		"Тест отрицательной цены",
		"Описание тестового объявления",
		"-100",
		"1",
		testPhotoPath,
	)

	t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
	require.Equal(t, http.StatusBadRequest, status, "При отрицательной цене ожидался статус 400")
	require.Equal(t, "price must be between 0 and 1000000 inclusive", apiError.Message, "При отрицательной цене сервер вернул неожиданное сообщение")
}

func TestCreateAdvertisementWithLongTitle(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуемся и создаём длинное название")
	login := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	longTitle := strings.Repeat("a", 1000)

	t.Log("Шаг: отправляем запрос с названием из 1000 символов")
	status, apiError := advertisements.CreateAdvertisementWithCleanup(
		t,
		login.Token,
		longTitle,
		"Описание тестового объявления",
		"100",
		"1",
		testPhotoPath,
	)

	t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
	require.Equal(t, http.StatusBadRequest, status, "При слишком длинном названии ожидался статус 400")
	require.Equal(t, "title must not exceed 50 characters provided 1000", apiError.Message, "При слишком длинном названии сервер вернул неожиданное сообщение")
}
func TestCreateAdvertisementWithoutPhoto(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуемся и готовим объявление без фотографии")
	login := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	t.Log("Шаг: отправляем запрос без фотографии")
	status, apiError := advertisements.CreateAdvertisementWithCleanup(
		t,
		login.Token,
		"Тестовое объявление без фотографии",
		"Описание тестового объявления",
		"100",
		"1",
		"",
	)

	t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
	require.Equal(t, http.StatusBadRequest, status, "Без фотографии ожидался статус 400")
	require.Equal(t, "at least 1 photo is required", apiError.Message, "При запросе без фотографии сервер вернул неожиданное сообщение")
}
