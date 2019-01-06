package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// CurrentAppointments returns a list of existing appointments.
// GET /appointments
func (app *Application) CurrentAppointments(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Ok: true, Data: app.db.Taken})
}

// AvailableAppointments returns a list of free slots for a given date.
// GET /appointments/:token/:date/free
func (app *Application) AvailableAppointments(c echo.Context) error {
	token := c.Param("token")

	if err := app.CheckToken(token); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	date := c.Param("date")

	data, err := app.db.GetAvailability(date)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseAvailability{Slots: data})
}

// CreateAppointment requests a slot to be reserved on the date and time for the patient name.
// POST /appointments/:token/:date/:time/:name
func (app *Application) CreateAppointment(c echo.Context) error {
	token := c.Param("token")

	if err := app.CheckToken(token); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	date := c.Param("date")
	time := c.Param("time")
	name := c.Param("name")

	uuid, err := app.db.CreateApointment(date, time, name)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseAppointment{UUID: uuid})
}

// DeleteAppointment deletes an appointment.
// DELETE /appointments/:token/:id
func (app *Application) DeleteAppointment(c echo.Context) error {
	token := c.Param("token")

	if err := app.CheckToken(token); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	id := c.Param("id")

	if err := app.db.DeleteAppointment(id); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Ok: true})
}
