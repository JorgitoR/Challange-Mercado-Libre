package entrypoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

	h.Router.HandleFunc("/v1/mercadoPago/creditApplication", h.CreditApplication).Methods("POST")
	h.Router.HandleFunc("/v1/mercadoPago/AllLoan", h.AllLoan).Methods("GET")
	h.Router.HandleFunc("/v1/mercadoPago/RegisterPaymentMade", h.RegisterPaymentMade).Methods("POST")
	h.Router.HandleFunc("/v1/mercadoPago/RaiseDebt", h.RaiseDebt).Methods("GET")
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
