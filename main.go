package main

import (
	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"sample/handler"
	"sample/migrations"
	"sample/store"
)

func main() {
	// Creating GoFr app
	app := gofr.New()

	// Running migrations - UP
	if err := migration.Migrate("remote-config-data", dbmigration.NewGorm(app.GORM()),
		migrations.All(), dbmigration.UP, app.Logger); err != nil {
		app.Logger.Fatalf("Error in running migrations: %v", err)
	}

	empStore := store.New()
	empHandler := handler.New(empStore)

	// Creating routes
	app.POST("/employee", empHandler.Create)
	app.GET("/employee/{id}", empHandler.GetByID)
	app.PUT("/employee/{id}", empHandler.Update)
	app.DELETE("/employee/{id}", empHandler.Delete)

	// Starting server
	app.Start()
}
