package advertisementSearch

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
	title     string
	beginning string
	middle    string
	end       string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	s.loginData = auth.Login(
		s.T(),
		os.Getenv("TEST_LOGIN"),
		os.Getenv("TEST_PASSWORD"),
	)
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// SetupTest создаёт объявление для каждого теста поиска.
func (s *TestSuite) SetupTest() {
	t := s.T()
	t.Log("Подготовка: создаём объявление для теста поиска")
	unique := fmt.Sprintf("%x", time.Now().UnixNano())
	s.beginning = "av" + unique
	s.middle = unique
	s.end = unique + "by"
	s.title = s.beginning + "by"

	advertisement := advList.CreateAdvertisement(
		t,
		s.loginData.Token,
		s.title,
		"Объявление для проверки поиска",
		"100",
		"1",
		"advertisement.jpg",
	)

	require.NotEmpty(t, advertisement.ID, "Сервер не вернул ID объявления для поиска")

	t.Cleanup(func() {
		t.Log("Очистка: удаляем объявление для теста поиска")
		advList.DeleteAdvertisement(t, s.loginData.Token, advertisement.ID)
	})
}

func (s *TestSuite) TestSearchByTitle() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, s.middle)

	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})

	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, s.title)
	})
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", s.beginning},
		{"середина", s.middle},
		{"конец", s.end},
		{"нижний регистр", strings.ToLower(s.middle)},
		{"верхний регистр", strings.ToUpper(s.middle)},
		{"смешанный регистр", utils.MixCase(s.middle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, s.title)
		})
	}
}
