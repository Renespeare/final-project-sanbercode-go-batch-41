package model

import "time"

type Comment struct {
	ID 						int64			`json:"id"`
	User_id 			int				`json:"user_id"`
	Article_id 		int				`json:"article_id"`
	Description 	string		`json:"description"`
	Created_at 		time.Time	`json:"created_at"`
}