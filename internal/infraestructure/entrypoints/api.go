package entrypoints

import (
	"encoding/json"
	"net/http"

	"github.com/JorgitoR/MercadoCredito/internal/domain/model"
)

// MercadoCreditoService - the interface for our comment service
type Service interface {
	CreditApplication(payload model.CreditApplication) (model.ResponseCreditApplication, error)
	AllLoan(model.RaiseDebt) (model.ResponseAllLoan, error)
	RegisterPaymentMade(model.RegisterPaymentMade) (model.ResponseRegisterPaymentMade, error)
	RaiseDebt() (model.ResponseRaiseDebt, error)
}

// CreditApplication -
func (h *Handler) CreditApplication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var creditApplication model.CreditApplication
	if err := json.NewDecoder(r.Body).Decode(&creditApplication); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	response, err := h.Service.CreditApplication(creditApplication)
	if err != nil {
		sendErrorResponse(w, "Failed to post new Credit Application", err)
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// AllLoan -
func (h *Handler) AllLoan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// RegisterPaymentMade -
func (h *Handler) RegisterPaymentMade(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var registerPaymentMade model.RegisterPaymentMade
	if err := json.NewDecoder(r.Body).Decode(&registerPaymentMade); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	response, err := h.Service.RegisterPaymentMade(registerPaymentMade)
	if err != nil {
		sendErrorResponse(w, "Failed to post new Credit Application", err)
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func (h *Handler) RaiseDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
