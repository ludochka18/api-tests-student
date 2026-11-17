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

func TestCannotAddOwnAdvertisementToFavorites(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуем владельца объявления")
	owner := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	t.Log("Подготовка: создаём объявление")
	advertisement := advertisements.CreateAdvertisement(
		t,
		owner.Token,
		"Тест собственного объявления",
		"Объявление для проверки избранного",
		"100",
		"1",
		"../../testdata/advertisement.jpg",
	)
	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления")

	t.Cleanup(func() {
		t.Log("Очистка: удаляем тестовое объявление")
		advertisements.DeleteAdvertisement(t, owner.Token, advertisement.ID)
	})

	t.Log("Шаг: пытаемся добавить своё объявление в избранное")
	status, apiError := favorites.AddFavoriteExpectingError(
		t,
		owner.Token,
		advertisement.ID,
	)

	t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
	require.Equal(t, http.StatusForbidden, status, "Для своего объявления ожидался статус 403")
	require.Equal(
		t,
		"Cannot add own advertisement to favorites",
		apiError.Message,
		"При добавлении своего объявления сервер вернул неожиданное сообщение",
	)
}
