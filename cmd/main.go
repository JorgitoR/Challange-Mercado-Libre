package main

import (
	"log"
	"net/http"

	"github.com/JorgitoR/MercadoCredito/internal/application/usecases"
	"github.com/JorgitoR/MercadoCredito/internal/domain"
	"github.com/JorgitoR/MercadoCredito/internal/infraestructure/database"
	"github.com/JorgitoR/MercadoCredito/internal/infraestructure/entrypoints"
	"github.com/JorgitoR/MercadoCredito/internal/usecases"
)

// App - the struct which contains information about our app
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {

	postgresClient, err := database.NewDatabase()
	if err != nil {
		return err
	}
	// Domain
	domain := domain.New(postgresClient)

	// UseCases
	service := usecases.NewService(domain, postgresClient)

	// Infraestructure -
	handler := entrypoints.NewAPIService(service)

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
		log.Fatal("Error starting up our REST API")
	}

	/*
		dynamoClient := dynamo.New()
		repository := adapters.NewDynamoAdapter(dynamoClient)
		domain := domain.New(repository)
		service := usecases.NewService(domain, repository)
		api := entrypoints.NewApi(service)
	*/

}
