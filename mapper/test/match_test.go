package test

import (
	"github.com/go-playground/assert/v2"
	"github.com/yakob-abada/restapi/mapper"
	"github.com/yakob-abada/restapi/model"
	"testing"
	"time"
)

func TestMatchMapper(t *testing.T) {
	t.Run("item", func(t *testing.T) {
		match := &model.Match{
			RecipientUserId: 1,
			ActorUserId:     2,
			Status:          model.MatchStatusUnMatched,
			CreatedAt:       time.Now(),
		}

		sut := mapper.NewMatchMapper()
		result := sut.Map(match)

		assert.Equal(t, match.RecipientUserId, result.MatchedUserId)
		assert.Equal(t, match.ActorUserId, result.UserId)
		assert.Equal(t, match.CreatedAt, result.CreatedAt)
	})

	t.Run("list", func(t *testing.T) {
		match := []*model.Match{
			{
				RecipientUserId: 1,
				ActorUserId:     2,
				Status:          model.MatchStatusUnMatched,
				CreatedAt:       time.Now(),
			},
		}

		sut := mapper.NewMatchMapper()
		result := sut.MapList(match)

		assert.Equal(t, match[0].RecipientUserId, result[0].MatchedUserId)
		assert.Equal(t, match[0].ActorUserId, result[0].UserId)
		assert.Equal(t, match[0].CreatedAt, result[0].CreatedAt)
	})
}
