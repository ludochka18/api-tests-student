package advertisements

import (
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"net/http"
	"testing"
)

// HttpGetAdvertisements отправляет запрос поиска и проверяет статус.
func HttpGetAdvertisements(t *testing.T, token, textForSearch string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementsPath).Query("search", textForSearch).Expect(t).Status(http.StatusOK).End().Response
}

// HttpCreateAdvertisement отправляет запрос на создание объявления с фотографией.
func HttpCreateAdvertisement(
	t *testing.T,
	token string,
	title string,
	description string,
	price string,
	quantity string,
	photoPath string,
	expectedStatus int,
) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Post("/advertisement").
		MultipartFormData("title", title). // MultipartFormData - добавляет текстовое поле в запрос
		MultipartFormData("description", description).
		MultipartFormData("price", price).
		MultipartFormData("quantity", quantity).
		MultipartFile("photos", photoPath).
		Expect(t).
		Status(expectedStatus).
		End().
		Response
}

// HttpDeleteAdvertisement удаляет объявление по его ID.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Delete("/advertisement").
		Query("id", id).
		Expect(t).
		Status(http.StatusNoContent).
		End().
		Response
}

// HttpGetAdvertisement получает объявление по его ID.
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Create().
		Get("/advertisement").
		Query("id", id).
		Expect(t).
		Status(http.StatusOK).
		End().
		Response
}

// HttpGetAdvertisementPhotos получает фотографии объявления по его ID.
func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()

	return apiRunner.GetRunner().
		Auth(token).
		Create().
		Get("/advertisements/" + id + "/photos").
		Expect(t).
		Status(http.StatusOK).
		End().
		Response
}

// HttpCreateAdvertisementWithoutStatusCheck отправляет запрос без проверки статуса.
func HttpCreateAdvertisementWithoutStatusCheck(
	t *testing.T,
	token string,
	title string,
	description string,
	price string,
	quantity string,
	photoPath string,
) *http.Response {
	t.Helper()

	request := apiRunner.GetRunner().
		Auth(token).
		Create().
		Post("/advertisement").
		MultipartFormData("title", title).
		MultipartFormData("description", description).
		MultipartFormData("price", price).
		MultipartFormData("quantity", quantity)

	if photoPath != "" {
		request.MultipartFile("photos", photoPath)
	}

	return request.
		Expect(t).
		End().
		Response
}
