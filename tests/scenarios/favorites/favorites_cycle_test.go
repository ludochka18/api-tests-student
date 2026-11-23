package favorites

import (
	"fmt"
	"testing"
	"time"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestAddAndRemoveFavorite(t *testing.T) {
	base.SetupSuite()

	owner, user := loginFavoriteUsers(t)

	title := fmt.Sprintf("Избранное %d", time.Now().UnixNano())

	t.Log("Подготовка: первый пользователь создаёт объявление")
	advertisement := advertisements.CreateAdvertisement(
		t,
		owner.Token,
		title,
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

	t.Log("Шаг: второй пользователь добавляет объявление в избранное")
	favorite := favorites.AddFavorite(t, user.Token, advertisement.ID)
	added := true

	t.Cleanup(func() {
		if added {
			t.Log("Очистка: удаляем запись из избранного")
			favorites.DeleteFavorite(t, user.Token, advertisement.ID)
		}
	})

	require.Equal(
		t,
		advertisement.ID,
		favorite.Advertisement.ID,
		"В избранное добавлено не то объявление",
	)

	t.Log("Проверка: объявление появилось в избранном")
	list := favorites.GetFavorites(t, user.Token, title)
	require.Len(t, list.Items, 1, "В избранном ожидалось одно тестовое объявление")
	require.Equal(t, advertisement.ID, list.Items[0].Advertisement.ID)

	t.Log("Шаг: удаляем объявление из избранного")
	favorites.DeleteFavorite(t, user.Token, advertisement.ID)
	added = false

	t.Log("Проверка: объявление исчезло из избранного")
	list = favorites.GetFavorites(t, user.Token, title)
	require.Empty(t, list.Items, "Объявление осталось в избранном после удаления")
}
