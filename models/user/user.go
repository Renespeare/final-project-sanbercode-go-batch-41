package model

import (
	// "database/sql"
	"time"
)

type Register struct {
	ID 					int64						`json:"id"`
	Name 				string					`json:"name"`
	Email 			string					`json:"email" binding:"required"`
	Password 		string					`json:"password" binding:"required"`
	Bio					*string					`json:"bio"`
	Created_at 	time.Time				`json:"created_at"`
	Updated_at 	*string					`json:"updated_at"`
}

type Login struct {
	Email 		string 	`json:"email" binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
}

type Credential struct {
	User_id 		int64		`json:"id"`
	Uuid 				string	`json:"uuid"`
}