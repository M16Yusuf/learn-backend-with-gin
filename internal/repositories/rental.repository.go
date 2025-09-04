package repositories

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/belajar-gin/internal/models"
)

type RentalRepository struct {
	db *pgxpool.Pool
}

func NewRentalRepository(db *pgxpool.Pool) *RentalRepository {
	return &RentalRepository{
		db: db,
	}
}

func (r *RentalRepository) GetRental(reqContext context.Context, offset, limit int) ([]models.Rental, error) {
	// query ke database
	sql := "SELECT * FROM rentals LIMIT $2 OFFSET $1"
	values := []any{offset, limit}
	rows, err := r.db.Query(reqContext, sql, values...)
	if err != nil {
		log.Println("internal server error : ", err.Error())
		return []models.Rental{}, err
	}
	defer rows.Close()

	// mengolah data / membaca rows/record
	var rentals []models.Rental
	for rows.Next() {
		var rental models.Rental
		if err := rows.Scan(&rental.Id, &rental.Image, &rental.Name, &rental.User_id, &rental.Created_at, &rental.Updated_at); err != nil {
			log.Println("Scan Error, ", err.Error())
			return []models.Rental{}, err
		}
		rentals = append(rentals, rental)
	}
	return rentals, nil
}

func (r *RentalRepository) NewRental(reqContext context.Context, body models.Rental) (models.Rental, error) {
	sql := "INSERT INTO rentals (image, rentals_name, user_id) VALUES ($1, $2, $3) RETURNING id, image, rentals_name, user_id, created_at"
	values := []any{body.Image, body.Name, body.User_id}
	var newRental models.Rental
	if err := r.db.QueryRow(reqContext, sql, values...).Scan(&newRental.Id, &newRental.Image, &newRental.Name, &newRental.User_id, &newRental.Created_at); err != nil {
		log.Println("Scan Error, ", err.Error())
		return models.Rental{}, err
	}
	return newRental, nil
}

func (r *RentalRepository) UpdateRental(reqContext context.Context, body models.Rental) (models.Rental, error) {
	// query dan pengecekan apa saja yang akan dirubah
	sql := "UPDATE rentals SET updated_at=CURRENT_TIMESTAMP "
	if body.Image != "" {
		sql += ", image='" + body.Image + "'"
	}
	if body.Name != "" {
		sql += ", rentals_name='" + body.Name + "'"
	}
	if body.User_id != 0 {
		sql += ", user_id=" + strconv.Itoa(body.User_id) + ""
	}
	sql += " WHERE id=$1 RETURNING id, image, rentals_name, user_id, created_at, updated_at "
	values := []any{body.Id}

	// masukan data hasil returning query ke dalam model rental yang baru
	var newRental models.Rental
	if err := r.db.QueryRow(reqContext, sql, values...).Scan(&newRental.Id, &newRental.Image, &newRental.Name, &newRental.User_id, &newRental.Created_at, &newRental.Updated_at); err != nil {
		log.Println("scan Error. ", err.Error())
		return models.Rental{}, err
	}

	return newRental, nil
}
