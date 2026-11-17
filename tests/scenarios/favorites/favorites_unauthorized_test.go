package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestFavoritesWithoutAuthorization(t *testing.T) {
	base.SetupSuite()

	const advertisementID = "00000000-0000-4000-8000-000000000001"
	const emptyToken = ""
	const expectedMessage = "Authorization header required"

	t.Run("Получение списка без токена", func(t *testing.T) {
		t.Log("Шаг: запрашиваем избранное без авторизации")
		status, apiError := favorites.GetFavoritesExpectingError(t, emptyToken)

		t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
		require.Equal(t, http.StatusUnauthorized, status, "Ожидался статус 401")
		require.Equal(t, expectedMessage, apiError.Message, "Неожиданное сообщение об авторизации")
	})

	t.Run("Добавление без токена", func(t *testing.T) {
		t.Log("Шаг: добавляем объявление без авторизации")
		status, apiError := favorites.AddFavoriteExpectingError(
			t, emptyToken, advertisementID,
		)

		t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
		require.Equal(t, http.StatusUnauthorized, status, "Ожидался статус 401")
		require.Equal(t, expectedMessage, apiError.Message, "Неожиданное сообщение об авторизации")
	})

	t.Run("Удаление без токена", func(t *testing.T) {
		t.Log("Шаг: удаляем объявление без авторизации")
		status, apiError := favorites.DeleteFavoriteExpectingError(
			t, emptyToken, advertisementID,
		)

		t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
		require.Equal(t, http.StatusUnauthorized, status, "Ожидался статус 401")
		require.Equal(t, expectedMessage, apiError.Message, "Неожиданное сообщение об авторизации")
	})
}
