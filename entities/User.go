package entities

import "time"


type User struct{
	Id string
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}