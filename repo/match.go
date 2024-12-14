package repo

import (
	"github.com/yakob-abada/restapi/model"
	"gorm.io/gorm"
)

type Match struct {
	db *gorm.DB
}

func NewMatch(db *gorm.DB) *Match {
	return &Match{
		db: db,
	}
}

func (m *Match) MatchWithRecipient(userId int, statuses []int, paginatedReq *PaginatedRequest) ([]*model.Match, error) {
	if paginatedReq == nil {
		paginatedReq = DefaultPaginatedRequest()
	}

	var matches []*model.Match
	err := m.db.Offset(paginatedReq.Offset()).Limit(paginatedReq.Limit()+1).
		Where("recipient_user_id = ? AND status in (?)", userId, statuses).
		Find(&matches).Error

	if err != nil {
		return nil, err
	}

	return matches, nil
}
