package postgres

import (
	"time"

	"github.com/nori-plugins/profile/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type:bigint; not null" `
	FirstName string    `gorm:"column:first_name; type:VARCHAR(25)"`
	LastName  string    `gorm:"column:last_name; type:VARCHAR(25)"`
	NickName  string    `gorm:"column:nick_name; type:VARCHAR(25)"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.Profile {
	return &entity.Profile{
		ID:        m.ID,
		UserID:    m.UserID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		NickName:  m.NickName,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func newModel(e *entity.Profile) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		NickName:  e.NickName,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "profiles"
}
