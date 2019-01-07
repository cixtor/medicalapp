package main

import "testing"

func TestCheckTokenValid(t *testing.T) {
	app := NewApp("c5e278ae")
	if err := app.CheckToken("c5e278ae"); err != nil {
		t.Fatal("token validation is incorrect")
	}
}

func TestCheckTokenInvalid(t *testing.T) {
	app := NewApp("c5e278ae")
	err := app.CheckToken("b8f7610f")
	if err.Error() != "token is invalid" {
		t.Fatal("token validation is incorrect")
	}
}

func TestCheckTokenMissing(t *testing.T) {
	app := NewApp("c5e278ae")
	err := app.CheckToken("")
	if err.Error() != "token is missing" {
		t.Fatal("token validation is incorrect")
	}
}

func TestGetAvailability(t *testing.T) {
	db := NewDatabase()
	slots, err := db.GetAvailability("2018-07-20")
	if err != nil {
		t.Fatal("availability getter failed:", err)
	}
	if len(slots) < 48 {
		t.Fatal("all slots should be available")
	}
}

func TestGetAvailabilityTaken(t *testing.T) {
	db := NewDatabase()
	db.CreateApointment("2018-07-20", "00:00", "John")
	db.CreateApointment("2018-07-20", "01:00", "John")
	db.CreateApointment("2018-07-20", "02:00", "John")
	db.CreateApointment("2018-07-20", "03:00", "John")
	db.CreateApointment("2018-07-20", "04:00", "John")
	db.CreateApointment("2018-07-20", "05:00", "John")
	db.CreateApointment("2018-07-20", "06:00", "John")
	db.CreateApointment("2018-07-20", "07:00", "John")
	db.CreateApointment("2018-07-20", "08:00", "John")
	db.CreateApointment("2018-07-20", "09:00", "John")
	db.CreateApointment("2018-07-20", "10:00", "John")
	slots, err := db.GetAvailability("2018-07-20")
	if err != nil {
		t.Fatal("availability getter failed:", err)
	}
	if len(slots) < 37 {
		t.Fatal("some slots should be taken")
	}
}

func TestGetAvailabilityInvalid(t *testing.T) {
	db := NewDatabase()
	_, err := db.GetAvailability("some-bad-date")
	if err.Error() != `parsing time "some-bad-date" as "2006-01-02": cannot parse "some-bad-date" as "2006"` {
		t.Fatal("invalid date check", err)
	}
}

func TestCreateApointment(t *testing.T) {
	db := NewDatabase()

	uuid, err := db.CreateApointment("2018-07-20", "10:30", "John")

	if err != nil {
		t.Fatal("something unexpected;", err)
	}

	if !db.TimeIsTaken("2018-07-20", "10:30") {
		t.Fatal("time slot should have been taken")
	}

	db.DeleteAppointment(uuid)

	if db.TimeIsTaken("2018-07-20", "10:30") {
		t.Fatal("time slot should have been released")
	}
}
