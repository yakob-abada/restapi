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

func TestLike(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		now := time.Now()

		matchListModel := []*model.Match{
			{
				RecipientUserId: 1,
				ActorUserId:     2,
				Status:          model.MatchStatusPending,
				CreatedAt:       now,
			},
		}
		likeRepoMock := &repo.MatchMock{}
		likeRepoMock.On("MatchWithRecipient", 1, []int{model.MatchStatusPending}, mock.Anything).Return(matchListModel, nil)
		authMock := &auth.AuthorizationMock{}
		authMock.On("UserId").Return(1)
		likeListMap := []*mapper.Like{
			{
				UserId:      2,
				LikedUserId: 1,
				CreatedAt:   now,
			},
		}

		likeMapMock := &mapper.LikeMock{}
		likeMapMock.On("MapList", mock.Anything).Return(likeListMap)

		sut := handler.NewLike(authMock, likeRepoMock, likeMapMock)
		sut.WhoLikedMe(c)

		responseData, _ := io.ReadAll(w.Body)
		var list []*mapper.Like
		json.Unmarshal(responseData, &list)

		assert.Equal(t, likeListMap[0].LikedUserId, list[0].LikedUserId)
		assert.Equal(t, likeListMap[0].UserId, list[0].UserId)
		assert.Equal(t, likeListMap[0].CreatedAt.Format("2006-01-02T15:04:05"), list[0].CreatedAt.Format("2006-01-02T15:04:05"))
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("repo fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		likeRepoMock := &repo.MatchMock{}
		likeRepoMock.On("MatchWithRecipient", 1, []int{model.MatchStatusPending}, mock.Anything).Return([]*model.Match{}, errors.New("failed connection"))
		authMock := &auth.AuthorizationMock{}
		authMock.On("UserId").Return(1)

		likeMapMock := &mapper.LikeMock{}

		sut := handler.NewLike(authMock, likeRepoMock, likeMapMock)
		sut.WhoLikedMe(c)

		responseData, _ := io.ReadAll(w.Body)
		var list []*mapper.Like
		json.Unmarshal(responseData, &list)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
