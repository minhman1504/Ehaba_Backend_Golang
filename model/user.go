package model

import "time"

type User struct {
	UserId string `json:"-" bson:"_id,omitempty"`

	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`

	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"` // tags
	Gender    string `json:"gender,omitempty"`
	Birthday  string `json:"birthday,omitempty"`

	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
	Token     string    `json:"token,omitempty"`
}
