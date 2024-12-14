package test

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/restapi/auth"
	"github.com/yakob-abada/restapi/handler"
	"github.com/yakob-abada/restapi/mapper"
	"github.com/yakob-abada/restapi/model"
	"github.com/yakob-abada/restapi/repo"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		modelProfile := []*model.Profile{
			{
				ID:       1,
				Name:     "test",
				Gender:   "male",
				Lat:      "",
				Long:     "",
				DietType: "vegan",
				Age:      11,
			},
		}
		profileRepoMock := &repo.ProfileMock{}
		profileRepoMock.On("Explore", 1, mock.Anything).Return(modelProfile, nil)
		authMock := &auth.AuthorizationMock{}
		authMock.On("UserId").Return(1)

		sut := handler.NewProfile(authMock, profileRepoMock)
		sut.Explore(c)

		responseData, _ := io.ReadAll(w.Body)
		var list []*model.Profile
		json.Unmarshal(responseData, &list)

		assert.Equal(t, modelProfile[0].ID, list[0].ID)
		assert.Equal(t, modelProfile[0].Name, list[0].Name)
		assert.Equal(t, modelProfile[0].Gender, list[0].Gender)
		assert.Equal(t, modelProfile[0].Lat, list[0].Lat)
		assert.Equal(t, modelProfile[0].Long, list[0].Long)
		assert.Equal(t, modelProfile[0].DietType, list[0].DietType)
		assert.Equal(t, modelProfile[0].Age, list[0].Age)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("repo fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		profileRepoMock := &repo.ProfileMock{}
		profileRepoMock.On("Explore", 1, mock.Anything).Return([]*model.Profile{}, errors.New("failed connection"))
		authMock := &auth.AuthorizationMock{}
		authMock.On("UserId").Return(1)

		sut := handler.NewProfile(authMock, profileRepoMock)
		sut.Explore(c)

		responseData, _ := io.ReadAll(w.Body)
		var list []*mapper.Like
		json.Unmarshal(responseData, &list)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
