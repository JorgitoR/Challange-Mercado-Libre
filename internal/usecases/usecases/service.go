package usecases

import (
	"github.com/JorgitoR/MercadoCredito/internal/domain/model"
	"github.com/jinzhu/gorm"
)

type DomainMarketPlace interface {
	CreditApplication(payload model.CreditApplication) (model.ResponseCreditApplication, error)
	AllLoan() (model.ResponseAllLoan, error)
	RegisterPaymentMade(model.RegisterPaymentMade) (model.ResponseRegisterPaymentMade, error)
	RaiseDebt(model.RaiseDebt) (model.ResponseRaiseDebt, error)
}

type Service struct {
	DB     *gorm.DB
	domain DomainMarketPlace
}

// NewService - returns a new Market Credito service
func NewService(domain DomainMarketPlace, db *gorm.DB) *Service {
	return &Service{
		domain: domain,
		DB:     db,
	}
}

// CreditApplication - adds a new Credit Application
func (s *Service) CreditApplication(payload model.CreditApplication) (model.ResponseCreditApplication, error) {
	response, err := s.domain.CreditApplication(payload)
	if err != nil {
		return model.ResponseCreditApplication{}, err
	}
	return response, nil
}

// AllLoan - retrieves all loan [date from, date to]
func (s *Service) AllLoan() (model.ResponseAllLoan, error) {
	response, err := s.domain.AllLoan()
	if err != nil {
		return model.ResponseAllLoan{}, nil
	}
	return response, nil
}

// RegisterPaymentMade - adds a new register of one pay realized
func (s *Service) RegisterPaymentMade(payload model.RegisterPaymentMade) (model.ResponseRegisterPaymentMade, error) {
	response, err := s.domain.RegisterPaymentMade(payload)
	if err != nil {
		return model.ResponseRegisterPaymentMade{}, nil
	}
	return response, nil
}

// RaiseDebt -
func (s *Service) RaiseDebt(payload model.RaiseDebt) (model.ResponseRaiseDebt, error) {
	response, err := s.domain.RaiseDebt(payload)
	if err != nil {
		return model.ResponseRaiseDebt{}, err
	}
	return response, nil
}
