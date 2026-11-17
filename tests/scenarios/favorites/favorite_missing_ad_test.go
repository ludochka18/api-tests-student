package favorites

import (
	"net/http"
	"os"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestFavoriteWithDeletedAdvertisement(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуем два аккаунта")
	owner := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	user := auth.Login(t, os.Getenv("TEST_SECOND_LOGIN"), os.Getenv("TEST_SECOND_PASSWORD"))
	require.NotEqual(t, owner.User.Id, user.User.Id, "Для теста нужны разные аккаунты")

	t.Log("Подготовка: создаём объявление")
	advertisement := advertisements.CreateAdvertisement(
		t,
		owner.Token,
		"Объявление для проверки ошибок",
		"Объявление будет удалено перед проверкой",
		"100",
		"1",
		"../../testdata/advertisement.jpg",
	)
	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления")

	deleted := false
	t.Cleanup(func() {
		if !deleted {
			t.Log("Очистка: удаляем тестовое объявление")
			advertisements.DeleteAdvertisement(t, owner.Token, advertisement.ID)
		}
	})

	t.Log("Подготовка: удаляем объявление")
	advertisements.DeleteAdvertisement(t, owner.Token, advertisement.ID)
	deleted = true

	t.Run("Добавление несуществующего объявления", func(t *testing.T) {
		t.Log("Шаг: добавляем удалённое объявление в избранное")
		status, apiError := favorites.AddFavoriteExpectingError(
			t,
			user.Token,
			advertisement.ID,
		)

		t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
		require.Equal(t, http.StatusNotFound, status, "Ожидался статус 404")
		require.Equal(
			t,
			"Advertisement not found",
			apiError.Message,
			"Неожиданное сообщение при добавлении",
		)
	})

	t.Run("Удаление несуществующего объявления", func(t *testing.T) {
		t.Log("Шаг: удаляем несуществующее объявление из избранного")
		status, apiError := favorites.DeleteFavoriteExpectingError(
			t,
			user.Token,
			advertisement.ID,
		)

		t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
		require.Equal(t, http.StatusNotFound, status, "Ожидался статус 404")
		require.Equal(
			t,
			"Favorite entry not found",
			apiError.Message,
			"Неожиданное сообщение при удалении",
		)
	})
}
