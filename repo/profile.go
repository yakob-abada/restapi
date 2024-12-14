package repo

import (
	"github.com/yakob-abada/restapi/model"
	"gorm.io/gorm"
)

type Profile struct {
	db *gorm.DB
}

func NewProfile(db *gorm.DB) *Profile {
	return &Profile{
		db: db,
	}
}

func (p *Profile) Explore(userId int, paginatedReq *PaginatedRequest) ([]*model.Profile, error) {
	if paginatedReq == nil {
		paginatedReq = DefaultPaginatedRequest()
	}

	var profiles []*model.Profile
	err := p.db.Offset(paginatedReq.Offset()).Limit(paginatedReq.Limit()+1).
		Joins(
			"JOIN explores ON profiles.gender = explores.gender AND "+
				"profiles.age BETWEEN explores.age_from AND explores.age_to AND "+
				"explores.diet_type = profiles.diet_type",
		).Joins("JOIN profiles as actor ON actor.id = ? AND "+
		"ST_Distance_Sphere("+
		"point(actor.lat, actor.long), point(profiles.lat, profiles.long)"+
		") <= explores.distance_radius", userId).
		Find(&profiles).Error

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
