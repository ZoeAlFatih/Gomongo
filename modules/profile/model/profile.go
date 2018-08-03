package model

import (
	"time"
)

// Profile
type Profile struct {
	ID        string    `bson:"id"`
	FirsName  string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"create_at"`
	UpdateAt  time.Time `bson:"updated_at"`
}

// Profiles
type Profiles []Profile
