package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
)

// foobar is the most secure token in the world; trust me :D
const foobar string = "2bf8d578-7ab4-4694-b5e6-d89b794872c1"

func main() {
	app := NewApp(foobar)

	e := echo.New()

	/* the application doesn't supports CORS */
	e.GET("/appointments", app.CurrentAppointments)
	e.GET("/appointments/:token/:date/free", app.AvailableAppointments)
	e.POST("/appointments/:token/:date/:time/:name", app.CreateAppointment)
	e.DELETE("/appointments/:token/:id", app.DeleteAppointment)

	go func() {
		if err := e.Start(":80"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit /* graceful shutdown; timeout in 10 secs */
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
