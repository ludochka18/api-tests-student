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
