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

// Explore is return a list of profiles that matches users criteria and hasn't been liked already by the user OR has been not matched nor unmatched to eachother.
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
		Joins("LEFT JOIN matches AS likers ON likers.actor_user_id = ? AND likers.recipient_user_id = profiles.id", userId).
		Joins("LEFT JOIN matches AS matched ON matched.actor_user_id = profiles.id AND matched.recipient_user_id = ? AND matched.status in (1, 2)", userId).
		Where("(likers.recipient_user_id != profiles.id OR likers.recipient_user_id IS NULL) AND (matched.actor_user_id != profiles.id OR matched.actor_user_id IS NULL)").
		Group("profiles.id").Find(&profiles).Error

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
