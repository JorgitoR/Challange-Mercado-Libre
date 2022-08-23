package domain

import (
	"github.com/JorgitoR/MercadoCredito/internal/domain/model"
	_ "github.com/JorgitoR/MercadoCredito/internal/domain/model"
)

type Repository interface {
	CreditApplication(payload model.CreditApplication) (model.ResponseCreditApplication, error)
	AllLoan() (model.ResponseAllLoan, error)
	RegisterPaymentMade(model.RegisterPaymentMade) (model.ResponseRegisterPaymentMade, error)
	RaiseDebt(model.RaiseDebt) (model.ResponseRaiseDebt, error)
}
