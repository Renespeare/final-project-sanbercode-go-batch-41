package model

import "time"

type Article struct {
	ID 						int64			`json:"id"`
	User_id 			int				`json:"user_id"`
	Category_id 	int				`json:"category_id"`
	Title 				string		`json:"title"`
	Description 	string		`json:"description"`
	Created_at 		time.Time	`json:"created_at"`
	Updated_at 		*string		`json:"updated_at"`
}