package domain

import (
	"context"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
	_ "github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
)

type MercadoCredito struct {
	repository Repository
}

func New(repository Repository) *MercadoCredito {
	return &MercadoCredito{
		repository: repository,
	}
}

func (m *MercadoCredito) CreditApplication(creditApplication model.CreditApplication) error {
	err := m.repository.CreditApplication(creditApplication)
	if err != nil {
		return err
	}
	return nil
}

func (m *MercadoCredito) UserLoans(creditApplication model.UserLoans) (uint, error) {
	loanId, err := m.repository.UserLoans(creditApplication)
	if err != nil {
		return 0, err
	}
	return loanId, nil
}

func (m *MercadoCredito) GetLoans(dateFrom string, dateTo string) ([]model.CreditApplication, error) {
	response, err := m.repository.GetLoans(dateFrom, dateTo)
	if err != nil {
		return []model.CreditApplication{}, err
	}
	return response, nil
}

func (m *MercadoCredito) PostPayment(payment model.RegisterPaymentMade) (model.DebtPayment, error) {
	response, err := m.repository.RegisterPaymentMade(payment)
	if err != nil {
		return model.DebtPayment{}, nil
	}
	return response, nil
}

func (m *MercadoCredito) GetDebt(id uint) (model.UserLoans, error) {

	response, err := m.repository.GetDebt(id)
	if err != nil {
		return model.UserLoans{}, err
	}
	return response, nil
}

///////////////////////////////

func (m *MercadoCredito) SavePayment(ctx context.Context, args model.DebtPayment) error {
	m.repository.SavePaymentDatabase(ctx, args)
	return nil
}
