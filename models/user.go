package models

import (
    "github.com/google/uuid"
    "time"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"`
    Name      string    `json:"name"`
    Email     string    `gorm:"unique" json:"email"`
    CreatedAt time.Time `json:"created_at"`
}