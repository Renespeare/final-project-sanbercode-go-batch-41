package structs

import (
	"time"
)

type Category struct {
	ID 					int64		`json:"id"`
	Name				string	`json:"name"`
	Created_at	time.Time	`json:"created_at"`
	Updated_at 	*string	`json:"updated_at"`
}