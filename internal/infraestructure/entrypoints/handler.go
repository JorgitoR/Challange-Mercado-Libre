package entrypoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	Router  *mux.Router
	Service Service
}

type Response struct {
	Message string
	Error   string
}

func NewAPIService(service Service) *Handler {
	return &Handler{Service: service}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/v1/credit", h.PostCredit).Methods("POST")
	h.Router.HandleFunc("/api/v1/credits", h.GetLoans).Methods("GET")
	h.Router.HandleFunc("/api/v1/payment", h.PostPayment).Methods("POST")
	h.Router.HandleFunc("/api/v1/debt", h.GetDebt).Methods("GET")

	h.Router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
