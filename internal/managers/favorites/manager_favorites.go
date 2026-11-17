package favorites

import (
	"encoding/json"
	"net/http"
	"testing"

	client "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/managers/favorites/models"

	"github.com/stretchr/testify/require"

	advertisementModels "api-tests-template/internal/managers/advertisements/models"
)

// GetFavorites возвращает первую страницу избранного с поиском по названию.
func GetFavorites(t *testing.T, token, search string) models.FavoriteList {
	t.Helper()

	response := client.HttpGetFavorites(t, token, search)
	defer response.Body.Close()
	require.Equal(t, http.StatusOK, response.StatusCode, "Не удалось получить избранное")

	var favorites models.FavoriteList
	err := json.NewDecoder(response.Body).Decode(&favorites)
	require.NoError(t, err, "Не удалось прочитать список избранного")

	return favorites
}

// AddFavorite добавляет объявление в избранное и возвращает созданную запись.
func AddFavorite(t *testing.T, token, advertisementID string) models.FavoriteItem {
	t.Helper()

	response := client.HttpAddFavorite(t, token, advertisementID)
	defer response.Body.Close()
	require.Equal(t, http.StatusCreated, response.StatusCode, "Не удалось добавить объявление %s в избранное", advertisementID)

	var favorite models.FavoriteItem
	err := json.NewDecoder(response.Body).Decode(&favorite)
	require.NoError(t, err, "Не удалось прочитать ответ добавления в избранное")

	return favorite
}

// DeleteFavorite удаляет объявление из избранного.
func DeleteFavorite(t *testing.T, token, advertisementID string) {
	t.Helper()

	response := client.HttpDeleteFavorite(t, token, advertisementID)
	defer response.Body.Close()
	require.Equal(t, http.StatusNoContent, response.StatusCode, "Не удалось удалить объявление %s из избранного", advertisementID)
}

// AddFavoriteExpectingError возвращает статус и ошибку добавления в избранное.
func AddFavoriteExpectingError(
	t *testing.T,
	token string,
	advertisementID string,
) (int, advertisementModels.ErrorResponse) {
	t.Helper()

	response := client.HttpAddFavoriteRaw(t, token, advertisementID)
	defer response.Body.Close()

	var apiError advertisementModels.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&apiError)
	require.NoErrorf(t, err, "Не удалось прочитать ошибку добавления в избранное (HTTP %d)", response.StatusCode)

	return response.StatusCode, apiError
}

// DeleteFavoriteExpectingError возвращает статус и ошибку удаления из избранного.
func DeleteFavoriteExpectingError(
	t *testing.T,
	token string,
	advertisementID string,
) (int, advertisementModels.ErrorResponse) {
	t.Helper()

	response := client.HttpDeleteFavoriteRaw(t, token, advertisementID)
	defer response.Body.Close()

	var apiError advertisementModels.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&apiError)
	require.NoErrorf(t, err, "Не удалось прочитать ошибку удаления из избранного (HTTP %d)", response.StatusCode)

	return response.StatusCode, apiError
}

// GetFavoritesExpectingError возвращает статус и ошибку получения избранного.
func GetFavoritesExpectingError(
	t *testing.T,
	token string,
) (int, advertisementModels.ErrorResponse) {
	t.Helper()

	response := client.HttpGetFavoritesRaw(t, token, "")
	defer response.Body.Close()

	var apiError advertisementModels.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&apiError)
	require.NoErrorf(t, err, "Не удалось прочитать ошибку получения избранного (HTTP %d)", response.StatusCode)

	return response.StatusCode, apiError
}
