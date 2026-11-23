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

func TestFavoriteDisappearsAfterAdvertisementDeletion(t *testing.T) {
	base.SetupSuite()

	owner, user := loginFavoriteUsers(t)

	title := fmt.Sprintf("Удаление избранного %d", time.Now().UnixNano())

	t.Log("Подготовка: создаём объявление")
	advertisement := advertisements.CreateAdvertisement(
		t,
		owner.Token,
		title,
		"Объявление для проверки удаления",
		testPrice,
		testQuantity,
		testPhotoPath,
	)
	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления")

	deleted := false
	t.Cleanup(func() {
		if !deleted {
			t.Log("Очистка: удаляем тестовое объявление")
			advertisements.DeleteAdvertisement(t, owner.Token, advertisement.ID)
		}
	})

	t.Log("Шаг: второй пользователь добавляет объявление в избранное")
	favorites.AddFavorite(t, user.Token, advertisement.ID)

	t.Log("Проверка: объявление появилось в избранном")
	list := favorites.GetFavorites(t, user.Token, title)
	require.Len(t, list.Items, 1, "Объявление не найдено в избранном")
	require.Equal(
		t,
		advertisement.ID,
		list.Items[0].Advertisement.ID,
		"ID объявления в избранном не совпадает с ID созданного объявления",
	)

	t.Log("Шаг: владелец удаляет объявление")
	advertisements.DeleteAdvertisement(t, owner.Token, advertisement.ID)
	deleted = true

	t.Log("Проверка: удалённое объявление исчезло из избранного")
	list = favorites.GetFavorites(t, user.Token, title)
	require.Empty(t, list.Items, "Удалённое объявление осталось в избранном")
}
