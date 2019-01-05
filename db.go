package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Database represents a persistent data storage, PostgreSQL if you want.
//
// Taken is a key-value data structure where the key is the date, the value is
// a list of appointments which are also represented with a key-value structure
// where the key is the time in either hour-o'clock or hour-30mins and the value
// is a boolean True meaning that the appointment is set, False if it has been
// cancelled.
type Database struct {
	Taken map[string]Appointment
}

// Appointment represents a single appointment with a patient.
type Appointment struct {
	Date string
	Time string
	Name string
}

// NewDatabase initializes the database and returns its pointer.
func NewDatabase() *Database {
	db := new(Database)
	db.Taken = make(map[string]Appointment, 0)
	return db
}

// GetAvailability returns a list of free slots for a given date.
func (db *Database) GetAvailability(date string) ([]string, error) {
	var step time.Time

	t, err := time.Parse(`2006-01-02`, date)

	if err != nil {
		return []string{}, err
	}

	var timeStr string

	data := make([]string, 0)

	for hour := 0; hour < 1440; hour += 30 {
		step = t.Add(time.Minute * time.Duration(hour))
		timeStr = step.Format(`15:04`)

		if db.TimeIsTaken(date, timeStr) {
			continue
		}

		data = append(data, timeStr)
	}

	return data, nil
}

// CreateApointment inserts an appointment into the database.
func (db *Database) CreateApointment(date string, time string, name string) (string, error) {
	uuid := uuid.New().String()

	if db.TimeIsTaken(date, time) {
		return "", errors.New("time is already taken")
	}

	db.Taken[uuid] = Appointment{
		Date: date,
		Time: time,
		Name: name,
	}

	return uuid, nil
}

// DeleteAppointment cancels an already defined appointment.
func (db *Database) DeleteAppointment(uuid string) error {
	if _, ok := db.Taken[uuid]; !ok {
		return errors.New("appointment does not exists")
	}

	delete(db.Taken, uuid)

	return nil
}

// TimeIsTaken checks if the time is already used in an appointment.
func (db *Database) TimeIsTaken(date string, timeStr string) bool {
	for _, appo := range db.Taken {
		if appo.Date == date && appo.Time == timeStr {
			return true
		}
	}

	return false
}
