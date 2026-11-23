package favorites

import (
	"os"
	"testing"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"

	"github.com/stretchr/testify/require"
)

// loginFavoriteUsers авторизует владельца и пользователя избранного.
func loginFavoriteUsers(t *testing.T) (models.LoginOkResponse, models.LoginOkResponse) {
	t.Helper()

	t.Log("Подготовка: авторизуем владельца и пользователя избранного")
	owner := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	user := auth.Login(t, os.Getenv("TEST_SECOND_LOGIN"), os.Getenv("TEST_SECOND_PASSWORD"))

	require.NotEmpty(t, owner.User.Id, "Не удалось получить ID владельца")
	require.NotEmpty(t, user.User.Id, "Не удалось получить ID пользователя избранного")
	require.NotEqual(t, owner.User.Id, user.User.Id, "Для теста нужны два разных аккаунта")

	return owner, user
}
