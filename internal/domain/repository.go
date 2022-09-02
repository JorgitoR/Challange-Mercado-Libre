package domain

import (
	"context"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
	_ "github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
)

type Repository interface {
	CreditApplication(payload model.CreditApplication) error
	UpdateCreditApplication(model.CreditApplication) error
	RegisterPaymentMade(model.RegisterPaymentMade) (model.DebtPayment, error)
	GetLoans(dateFrom string, dateTo string) ([]model.CreditApplication, error)
	GetDebt(id uint) (model.UserLoans, error)
	UserLoans(model.UserLoans) (uint, error)
	SavePaymentDatabase(context.Context, model.DebtPayment) error
}
