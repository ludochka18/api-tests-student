package favorites

import (
	"os"
	"testing"

	"api-tests-template/internal/managers/auth"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestFavoriteAccounts(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуем первый аккаунт")
	first := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	t.Log("Подготовка: авторизуем второй аккаунт")
	second := auth.Login(t, os.Getenv("TEST_SECOND_LOGIN"), os.Getenv("TEST_SECOND_PASSWORD"))

	t.Log("Проверка: аккаунты принадлежат разным пользователям")
	require.NotEmpty(t, first.User.Id, "Не удалось получить ID первого пользователя")
	require.NotEmpty(t, second.User.Id, "Не удалось получить ID второго пользователя")
	require.NotEqual(t, first.User.Id, second.User.Id, "Для теста нужны два разных аккаунта")
}
