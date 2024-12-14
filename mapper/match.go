package mapper

import (
	"github.com/yakob-abada/restapi/model"
	"time"
)

type Match struct {
	UserId        int       `json:"user_id"`
	MatchedUserId int       `json:"matched_user_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type MatchMapper struct {
}

func NewMatchMapper() *MatchMapper {
	return &MatchMapper{}
}

func (m *MatchMapper) MapList(matches []*model.Match) []*Match {
	var list []*Match

	for _, match := range matches {
		list = append(list, m.Map(match))
	}

	return list
}

func (m *MatchMapper) Map(match *model.Match) *Match {
	return &Match{
		UserId:        match.ActorUserId,
		MatchedUserId: match.RecipientUserId,
		CreatedAt:     match.CreatedAt,
	}
}
