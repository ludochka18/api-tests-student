package advertisementCreate

import (
	"fmt"
	"os"
	"testing"
	"time"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
)

func TestCreateAdvertisement(t *testing.T) {
	base.SetupSuite()

	t.Log("Подготовка: авторизуемся и готовим данные объявления")
	login := auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))

	title := fmt.Sprintf("Тестовое объявление %d", time.Now().UnixNano())
	photoPath := testPhotoPath

	t.Log("Шаг: создаём объявление с фотографией")
	advertisement := advertisements.CreateAdvertisement(
		t,
		login.Token,
		title,
		"Описание тестового объявления",
		"100",
		"1",
		photoPath,
	)

	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления")

	t.Cleanup(func() {
		t.Log("Очистка: удаляем тестовое объявление")
		advertisements.DeleteAdvertisement(t, login.Token, advertisement.ID)
	})

	t.Log("Проверка: название в ответе совпадает с переданным")
	require.Equal(t, title, advertisement.Title, "Название объявления не совпадает")

	t.Log("Шаг: получаем созданное объявление по ID")
	savedAdvertisement := advertisements.GetAdvertisement(t, advertisement.ID)

	t.Log("Проверка: сервер сохранил данные объявления")
	require.Equal(t, advertisement.ID, savedAdvertisement.ID, "ID объявления не совпадает")
	require.Equal(t, title, savedAdvertisement.Title, "Название объявления не совпадает")
	require.Equal(t, "Описание тестового объявления", savedAdvertisement.Description, "Описание объявления не совпадает")
	require.Equal(t, float64(100), savedAdvertisement.Price, "Цена объявления не совпадает")
	require.Equal(t, 1, savedAdvertisement.Quantity, "Количество не совпадает")

	t.Log("Шаг: ищем объявление по полному названию")
	searchResult := advertisements.GetAdvertisementsList(t, login.Token, title)

	t.Log("Проверка: созданное объявление есть в результатах поиска")
	found := false
	for _, item := range searchResult.Items {
		if item.ID == advertisement.ID && item.Title == title {
			found = true
			break
		}
	}
	require.True(t, found, "Созданное объявление не найдено в поиске по полному названию")

	t.Log("Шаг: получаем фотографии созданного объявления")
	photos := advertisements.GetAdvertisementPhotos(t, login.Token, advertisement.ID)

	t.Log("Проверка: загруженная фотография доступна")
	require.Len(t, photos, 1, "Ожидалась одна фотография объявления")
	require.NotEmpty(t, photos[0].ID, "У фотографии отсутствует ID")
	require.NotEmpty(t, photos[0].URL, "У фотографии отсутствует URL")

	require.Len(t, advertisement.Photos, 1, "В ответе создания ожидалась одна фотография")
	require.Equal(
		t,
		advertisement.Photos[0].ID,
		photos[0].ID,
		"ID фотографии после получения не совпадает с ID при создании",
	)
}
