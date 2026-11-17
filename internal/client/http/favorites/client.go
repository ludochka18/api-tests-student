package favorites

import (
	"net/http"
	"testing"

	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpGetFavorites получает список избранного с поиском по названию.
func HttpGetFavorites(t *testing.T, token, search string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Get("/favorites").
		Query("search", search).
		Expect(t).
		Status(http.StatusOK).
		End().
		Response
}

// HttpAddFavorite добавляет объявление в избранное и проверяет статус ответа.
func HttpAddFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()

	response := HttpAddFavoriteRaw(t, token, advertisementID)

	if response.StatusCode != http.StatusCreated {
		t.Errorf(
			"Неожиданный статус добавления в избранное: ожидали %d, получили %d",
			http.StatusCreated,
			response.StatusCode,
		)
	}

	return response
}

// HttpDeleteFavorite удаляет объявление из избранного и проверяет статус ответа.
func HttpDeleteFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()

	response := HttpDeleteFavoriteRaw(t, token, advertisementID)

	if response.StatusCode != http.StatusNoContent {
		t.Errorf(
			"Неожиданный статус удаления из избранного: ожидали %d, получили %d",
			http.StatusNoContent,
			response.StatusCode,
		)
	}

	return response
}

// HttpAddFavoriteRaw отправляет запрос добавления и возвращает ответ без проверки статуса.
func HttpAddFavoriteRaw(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()

	request := map[string]string{
		"advertisement_id": advertisementID,
	}

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Post("/favorites").
		JSON(request).
		Expect(t).
		End().
		Response
}

// HttpDeleteFavoriteRaw отправляет запрос удаления и возвращает ответ без проверки статуса.
func HttpDeleteFavoriteRaw(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Delete("/favorites/" + advertisementID).
		Expect(t).
		End().
		Response
}

// HttpGetFavoritesRaw возвращает ответ списка избранного без проверки статуса.
func HttpGetFavoritesRaw(t *testing.T, token string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Get("/favorites").
		Expect(t).
		End().
		Response
}
