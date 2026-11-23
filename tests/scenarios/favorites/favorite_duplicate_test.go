package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestCannotAddFavoriteTwice(t *testing.T) {
	base.SetupSuite()

	owner, user := loginFavoriteUsers(t)

	t.Log("Подготовка: создаём объявление")
	advertisement := advertisements.CreateAdvertisement(
		t,
		owner.Token,
		"Тест повторного добавления",
		"Объявление для проверки избранного",
		testPrice,
		testQuantity,
		testPhotoPath,
	)
	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления")

	t.Cleanup(func() {
		t.Log("Очистка: удаляем тестовое объявление")
		advertisements.DeleteAdvertisement(t, owner.Token, advertisement.ID)
	})

	t.Log("Подготовка: добавляем объявление в избранное")
	favorites.AddFavorite(t, user.Token, advertisement.ID)

	t.Cleanup(func() {
		t.Log("Очистка: удаляем запись из избранного")
		favorites.DeleteFavorite(t, user.Token, advertisement.ID)
	})

	t.Log("Шаг: повторно добавляем то же объявление")
	status, apiError := favorites.AddFavoriteExpectingError(
		t,
		user.Token,
		advertisement.ID,
	)

	t.Logf("Результат: статус %d, сообщение: %s", status, apiError.Message)
	require.Equal(t, http.StatusConflict, status, "При повторном добавлении ожидался статус 409")
	require.Equal(
		t,
		"Advertisement is already in favorites",
		apiError.Message,
		"При повторном добавлении сервер вернул неожиданное сообщение",
	)
}
