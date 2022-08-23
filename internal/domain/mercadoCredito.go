package domain

import (
	"github.com/JorgitoR/MercadoCredito/internal/domain/model"
	_ "github.com/JorgitoR/MercadoCredito/internal/domain/model"
)

type MercadoCredito struct {
	repository Repository
}

func New(repository Repository) *MercadoCredito {
	return &MercadoCredito{
		repository: repository,
	}
}

func (m *MercadoCredito) CreditApplication(payload model.CreditApplication) (model.ResponseCreditApplication, error) {
	response, err := m.repository.CreditApplication(payload)
	if err != nil {
		return model.ResponseCreditApplication{}, err
	}
	return response, nil
}

func (m *MercadoCredito) AllLoan() (model.ResponseAllLoan, error) {
	response, err := m.repository.AllLoan()
	if err != nil {
		return model.ResponseAllLoan{}, err
	}
	return response, nil
}

func (m *MercadoCredito) RegisterPaymentMade(payload model.RegisterPaymentMade) (model.ResponseRegisterPaymentMade, error) {
	response, err := m.repository.RegisterPaymentMade(payload)
	if err != nil {
		return model.ResponseRegisterPaymentMade{}, nil
	}
	return response, nil
}

func (m *MercadoCredito) RaiseDebt(payload model.RaiseDebt) (model.ResponseRaiseDebt, error) {

	response, err := m.repository.RaiseDebt(payload)
	if err != nil {
		return model.ResponseRaiseDebt{}, nil
	}
	return response, nil
}
