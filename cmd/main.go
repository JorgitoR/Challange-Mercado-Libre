package main

import (
	"log"
	"net/http"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain"
	"github.com/JorgitoR/Challange-Mercado-Libre/internal/infraestructure/adapters"
	"github.com/JorgitoR/Challange-Mercado-Libre/internal/infraestructure/entrypoints"
	"github.com/JorgitoR/Challange-Mercado-Libre/internal/usecases"
)

// App - the struct which contains information about our app
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {

	postgresClient, err := adapters.PostgresClient()
	if err != nil {
		return err
	}
	err = adapters.MigrateDB(postgresClient)
	if err != nil {
		log.Fatal("failed to setup database.")
		return err
	}
	repository := struct {
		*adapters.DTBAdapter
	}{
		adapters.NewPostgreSQLAdapter(postgresClient),
	}
	// Domain
	domain := domain.New(repository)

	// UseCases
	domainMarketPlace := usecases.NewService(domain, postgresClient)

	// Infraestructure -
	handler := entrypoints.NewAPIService(domainMarketPlace)
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Fatal("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name:    "",
		Version: "1.0",
	}
	if err := app.Run(); err != nil {
		log.Fatal("Error starting up our REST API", err)
	}

}
