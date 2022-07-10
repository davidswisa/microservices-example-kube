package postgres

import (
	"database/sql"
	"fmt"

	"github.com/davidswisa/microservices-example-kube/pkg/reservation"

	_ "github.com/lib/pq"
)

type PGInfo struct {
	Port int    `json:"port"`
	DB   string `json:"db"`
	Host string `json:"host"`
	Pass string `json:"pass"`
	User string `json:"user"`
	SSL  string `json:"ssl"`
}

const (
	pgHost     string = "db"
	pgPort     int    = 5432
	pgUser     string = "postgres"
	pgPassword string = "postgres"
	pgDatabase string = "postgres"
	pgSslMode  string = "disable"
)

const (
	sqlInsert = `INSERT INTO reservations (id, name, date, party, hour) VALUES($1, $2, $3, $4, $5)`
	sqlUpdate = `UPDATE reservations SET name = $2, date = $3, party = $4, hour = $5 WHERE id=$1`
	sqlDelete = `DELETE FROM reservations WHERE id=$1`
	sqlSelect = `SELECT id, name, date, party, hour from reservations;`
)

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
	ssl      string
)

// Instance ...
type Instance struct {
	info string
	db   *sql.DB
}

// New ...
func New() (*Instance, error) {
	i := Instance{}
	if err := i.discover(); err != nil {
		return nil, err
	}

	if err := i.connect(); err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *Instance) discover() error {

	pgInfo := PGInfo{
		DB:   pgDatabase,
		Host: pgHost,
		Pass: pgPassword,
		Port: pgPort,
		SSL:  pgSslMode,
		User: pgUser,
	}

	i.info = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pgInfo.Host, pgInfo.Port, pgInfo.User, pgInfo.Pass, pgInfo.DB, pgInfo.SSL)

	return nil
}

func (i *Instance) connect() error {
	var err error
	i.db, err = sql.Open("postgres", i.info)
	return err
}

// Close the connection
func (i *Instance) Close() error {
	return i.db.Close()
}

// CreateReservation inserts a new reservation to postgres
func (i *Instance) CreateReservation(r reservation.Reservation) error {
	_, err := i.db.Exec(sqlInsert, r.ID, r.Name, r.Date, r.Party, r.Hour)
	return err
}

// UpdateReservation inserts a new reservation to postgres
func (i *Instance) UpdateReservation(r reservation.Reservation) error {
	_, err := i.db.Exec(sqlUpdate, r.ID, r.Name, r.Date, r.Party, r.Hour)
	return err
}

// DeleteReservation inserts a new reservation to postgres
func (i *Instance) DeleteReservation(id string) error {
	_, err := i.db.Exec(sqlDelete, id)
	return err
}

// AllReservations fetches all reservations from postgres
func (i *Instance) AllReservations() ([]reservation.Reservation, error) {
	var r []reservation.Reservation

	rows, err := i.db.Query(sqlSelect)

	defer rows.Close()
	for rows.Next() {
		rr := reservation.Reservation{}
		err = rows.Scan(&rr.ID, &rr.Name, &rr.Date, &rr.Party, &rr.Hour)
		if err != nil {
			panic(err)
		}
		r = append(r, rr)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
	return r, err
}
