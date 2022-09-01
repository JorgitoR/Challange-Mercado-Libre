package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		return fmt.Errorf("Failed to setup our database umm %+v ", err)
	}
	errDataMigrate := adapters.MigrateDB(postgresClient)
	if errDataMigrate != nil {
		log.Fatal("failed to setup database")
		return errDataMigrate
	}
	repository := struct {
		*adapters.DTBAdapter
	}{
		adapters.NewPostgreSQLAdapter(postgresClient),
	}
	// Domain
	domain := domain.New(repository)

	// UseCases
	domainMarketPlace := usecases.NewService(domain)

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
	fmt.Println(os.Getenv("DB_HOST"))
	app := App{
		Name:    "",
		Version: "1.0",
	}
	if err := app.Run(); err != nil {
		log.Fatal("Error starting up our REST API ", err)
	}

}
