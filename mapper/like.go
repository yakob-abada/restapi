package mapper

import (
	"github.com/yakob-abada/restapi/model"
	"time"
)

type Like struct {
	UserId      int       `json:"user_id"`
	LikedUserId int       `json:"liked_user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type LikeMapper struct{}

func NewLikeMapper() *LikeMapper {
	return &LikeMapper{}
}

func (m *LikeMapper) MapList(matches []*model.Match) []*Like {
	var list []*Like

	for _, match := range matches {
		list = append(list, m.Map(match))
	}

	return list
}

func (m *LikeMapper) Map(match *model.Match) *Like {
	return &Like{
		UserId:      match.ActorUserId,
		LikedUserId: match.RecipientUserId,
		CreatedAt:   match.CreatedAt,
	}
}
