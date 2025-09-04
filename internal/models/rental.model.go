package models

import "time"

type Rental struct {
	Id         int        `db:"id" json:"id"`
	Image      string     `db:"image" json:"image_path"`
	Name       string     `db:"rentals_name" json:"rental_name"`
	User_id    int        `db:"user_id" json:"user_id"`
	Created_at time.Time  `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
