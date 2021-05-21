package dbrepo

import (
	"context"
	"time"

	"github.com/NamanBalaji/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// Inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	statement := `INSERT INTO reservation (first_name, last_name, email, phone, start_date, end_date, room_id, created__at, updated_at)
				  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, statement,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
			     VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, statement, r.StartDate, r.EndDate, r.RoomID, r.ReservationID, time.Now(), time.Now(), r.RestrictionID)
	if err != nil {
		return err
	}
	return nil
}
