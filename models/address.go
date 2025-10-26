package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Street     string         `json:"street"`
	City       string         `json:"city"`
	State      string         `json:"state"`
	Country    string         `json:"country"`
	Number     string         `json:"number"`
	Complement string         `json:"complement"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (address *Address) BeforeCreate(tx *gorm.DB) (err error) {
	address.ID = uuid.New()
	return
}
