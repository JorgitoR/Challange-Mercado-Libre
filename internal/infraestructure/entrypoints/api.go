package entrypoints

import (
	"encoding/json"
	"net/http"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
)

// MercadoCreditoService - the interface for our comment service
type Service interface {
	PostCredit(payload model.CreditApplication) (model.ResponseCreditApplication, error)
	GetLoans(dateFrom string, dateTo string) ([]model.CreditApplication, error)
	PostPayment(model.RegisterPaymentMade) (model.DebtPayment, error)
	GetDebt(date string, target string) (model.Balnace, error)
}

// PostCredit -
func (h *Handler) PostCredit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var creditApplication model.CreditApplication
	if err := json.NewDecoder(r.Body).Decode(&creditApplication); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	response, err := h.Service.PostCredit(creditApplication)
	if err != nil {
		sendErrorResponse(w, "Failed to post new Credit Application", err)
		return
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// GetLoans -
func (h *Handler) GetLoans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dateFrom := r.URL.Query().Get("dateFrom")
	dateTo := r.URL.Query().Get("dateTo")

	response, err := h.Service.GetLoans(dateFrom, dateTo)
	if err != nil {
		sendErrorResponse(w, "Error Retrieving Loans By Date", err)
		return
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// RegisterPaymentMade -
func (h *Handler) PostPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var registerPaymentMade model.RegisterPaymentMade
	if err := json.NewDecoder(r.Body).Decode(&registerPaymentMade); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	response, err := h.Service.PostPayment(registerPaymentMade)
	if err != nil {
		sendErrorResponse(w, "Failed to Pay a Credit", err)
		return
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func (h *Handler) GetDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	date := r.URL.Query().Get("date")
	target := r.URL.Query().Get("target")

	response, err := h.Service.GetDebt(date, target)
	if err != nil {
		sendErrorResponse(w, "Error Retrieving Loans By Date", err)
		return
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}

}
