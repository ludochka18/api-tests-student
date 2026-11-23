package advertisements

import (
	client "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// GetAdvertisementsList возвращает результаты поиска по названию.
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	t.Helper()
	raw := client.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать ответ поиска объявлений")
	var result models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &result), "Не удалось разобрать JSON ответа поиска объявлений")
	return result
}

// CreateAdvertisement создаёт объявление и возвращает данные из ответа API.
func CreateAdvertisement(
	t *testing.T,
	token string,
	title string,
	description string,
	price string,
	quantity string,
	photoPath string,
) models.Advertisement {
	t.Helper()

	response := client.HttpCreateAdvertisement(
		t,
		token,
		title,
		description,
		price,
		quantity,
		photoPath,
		http.StatusCreated,
	)
	defer response.Body.Close()

	var advertisement models.Advertisement
	err := json.NewDecoder(response.Body).Decode(&advertisement)
	require.NoError(t, err, "Не удалось прочитать ответ создания объявления")

	return advertisement
}

// DeleteAdvertisement удаляет объявление по его ID.
func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()

	response := client.HttpDeleteAdvertisement(t, token, id)
	defer response.Body.Close()
	require.Equal(t, http.StatusNoContent, response.StatusCode, "Не удалось удалить объявление %s", id)
}

// GetAdvertisement получает данные объявления по его ID.
func GetAdvertisement(t *testing.T, id string) models.Advertisement {
	t.Helper()

	response := client.HttpGetAdvertisement(t, id)
	defer response.Body.Close()
	require.Equal(t, http.StatusOK, response.StatusCode, "Не удалось получить объявление %s", id)

	var advertisement models.Advertisement
	err := json.NewDecoder(response.Body).Decode(&advertisement)
	require.NoError(t, err, "Не удалось прочитать ответ получения объявления")

	return advertisement
}

// GetAdvertisementPhotos возвращает фотографии объявления по его ID.
func GetAdvertisementPhotos(t *testing.T, token, id string) []models.Photo {
	t.Helper()

	response := client.HttpGetAdvertisementPhotos(t, token, id)
	defer response.Body.Close()

	var photos []models.Photo
	err := json.NewDecoder(response.Body).Decode(&photos)
	require.NoError(t, err, "Не удалось прочитать список фотографий объявления")

	return photos
}

// CreateAdvertisementExpectingError отправляет запрос и возвращает ошибку API.
func CreateAdvertisementExpectingError(
	t *testing.T,
	token string,
	title string,
	description string,
	price string,
	quantity string,
	photoPath string,
	expectedStatus int,
) models.ErrorResponse {
	t.Helper()

	status, apiError := CreateAdvertisementWithCleanup(
		t, token, title, description, price, quantity, photoPath,
	)
	require.Equal(t, expectedStatus, status, "Неожиданный статус создания объявления")
	return apiError
}

// CreateAdvertisementWithCleanup отправляет запрос и удаляет неожиданно созданное объявление.
func CreateAdvertisementWithCleanup(
	t *testing.T,
	token string,
	title string,
	description string,
	price string,
	quantity string,
	photoPath string,
	extraPhotoPaths ...string,
) (int, models.ErrorResponse) {
	t.Helper()

	response := client.HttpCreateAdvertisementWithoutStatusCheck(
		t,
		token,
		title,
		description,
		price,
		quantity,
		photoPath,
		extraPhotoPaths...,
	)
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		var advertisement models.Advertisement
		err := json.NewDecoder(response.Body).Decode(&advertisement)
		require.NoError(t, err, "Не удалось прочитать созданное объявление")
		require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID созданного объявления")

		t.Cleanup(func() {
			t.Log("Очистка: удаляем неожиданно созданное объявление")
			DeleteAdvertisement(t, token, advertisement.ID)
		})

		return response.StatusCode, models.ErrorResponse{}
	}

	var apiError models.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&apiError)
	require.NoErrorf(t, err, "Не удалось прочитать ответ с ошибкой (HTTP %d)", response.StatusCode)

	return response.StatusCode, apiError
}
