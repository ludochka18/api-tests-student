package advertisementCreate

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestCreateAdvertisementWithoutAuthorization(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: готовим данные объявления без токена")

	t.Log("Шаг: отправляем запрос на создание без авторизации")
	apiError := advertisements.CreateAdvertisementExpectingError(
		t,
		"",
		"Тестовое объявление",
		"Описание тестового объявления",
		"100",
		"1",
		"../../testdata/advertisement.jpg",
		http.StatusUnauthorized,
	)

	t.Log("Проверка: сервер сообщает об отсутствии авторизации")
	require.Equal(t, "unauthorized", apiError.Error, "Не совпадает тип ошибки")
	require.Equal(t, "Authorization header required", apiError.Message, "Не совпадает сообщение об ошибке")
}

func TestCreateAdvertisementWithInvalidToken(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: готовим данные объявления с неверным токеном")

	t.Log("Шаг: отправляем запрос на создание с неверным токеном")
	apiError := advertisements.CreateAdvertisementExpectingError(
		t,
		"invalid-token",
		"Тестовое объявление",
		"Описание тестового объявления",
		"100",
		"1",
		"../../testdata/advertisement.jpg",
		http.StatusUnauthorized,
	)

	t.Log("Проверка: сервер вернул сообщение об ошибке авторизации")
	require.NotEmpty(t, apiError.Message, "Сервер не вернул сообщение об ошибке")
}
