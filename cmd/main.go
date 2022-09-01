package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// UseCases
	domainMarketPlace := usecases.NewService()

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
