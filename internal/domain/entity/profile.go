package entity

import "time"

type Profile struct {
	ID        uint64
	UserID    uint64
	FirstName string
	LastName  string
	NickName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
