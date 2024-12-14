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
	"time"
)

func TestMatchLike(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		now := time.Now()

		matchListModel := []*model.Match{
			{
				RecipientUserId: 1,
				ActorUserId:     2,
				Status:          model.MatchStatusMatched,
				CreatedAt:       now,
			},
		}
		matchListModelRepoMock := &repo.MatchMock{}
		matchListModelRepoMock.On("MatchWithRecipient", 1, []int{model.MatchStatusMatched}, mock.Anything).Return(matchListModel, nil)
		authMock := &auth.AuthorizationMock{}
		authMock.On("UserId").Return(1)
		matchListMap := []*mapper.Match{
			{
				UserId:        2,
				MatchedUserId: 1,
				CreatedAt:     now,
			},
		}

		matchMapMock := &mapper.MatchMock{}
		matchMapMock.On("MapList", mock.Anything).Return(matchListMap)

		sut := handler.NewMatch(authMock, matchListModelRepoMock, matchMapMock)
		sut.WeMatched(c)

		responseData, _ := io.ReadAll(w.Body)
		var list []*mapper.Match
		json.Unmarshal(responseData, &list)

		assert.Equal(t, matchListMap[0].MatchedUserId, list[0].MatchedUserId)
		assert.Equal(t, matchListMap[0].UserId, list[0].UserId)
		assert.Equal(t, matchListMap[0].CreatedAt.Format("2006-01-02T15:04:05"), list[0].CreatedAt.Format("2006-01-02T15:04:05"))
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("repo fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		matchListModelRepoMock := &repo.MatchMock{}
		matchListModelRepoMock.On("MatchWithRecipient", 1, []int{model.MatchStatusMatched}, mock.Anything).Return([]*model.Match{}, errors.New("failed connection"))
		authMock := &auth.AuthorizationMock{}
		authMock.On("UserId").Return(1)

		matchMapMock := &mapper.MatchMock{}

		sut := handler.NewMatch(authMock, matchListModelRepoMock, matchMapMock)
		sut.WeMatched(c)

		responseData, _ := io.ReadAll(w.Body)
		var list []*mapper.Like
		json.Unmarshal(responseData, &list)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
