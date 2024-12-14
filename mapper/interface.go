package mapper

import "github.com/yakob-abada/restapi/model"

type LikeMap interface {
	MapList([]*model.Match) []*Like
	Map(*model.Match) *Like
}

type MatchMap interface {
	MapList([]*model.Match) []*Match
	Map(*model.Match) *Match
}
