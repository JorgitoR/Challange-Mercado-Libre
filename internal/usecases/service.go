package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
	"github.com/jinzhu/gorm"
)

type DomainMarketPlace interface {
	CreditApplication(payload model.CreditApplication) error
	GetLoans(dateFrom string, dateTo string) ([]model.CreditApplication, error)
	PostPayment(model.RegisterPaymentMade) (model.DebtPayment, error)
	GetDebt(id uint) (model.UserLoans, error)
	UserLoans(model.UserLoans) (uint, error)
	SavePayment(context.Context, model.DebtPayment) error
}

type Service struct {
	DB     *gorm.DB
	DBSQL  *sql.DB
	domain DomainMarketPlace
}

// NewService - returns a new Market Credito service
func NewService(domain DomainMarketPlace, dbp *gorm.DB, db *sql.DB) *Service {
	return &Service{
		domain: domain,
		DB:     dbp,
		DBSQL:  db,
	}
}

// PostCredit - adds a new Credit Application
func (s *Service) PostCredit(creditApplication model.CreditApplication) (model.ResponseCreditApplication, error) {

	historyCredit, err := s.HistoryCredit(creditApplication)
	if err != nil {
		return model.ResponseCreditApplication{}, err
	}
	credit, err := s.FecadeTargetClient(historyCredit, creditApplication)
	if err != nil {
		return model.ResponseCreditApplication{}, err
	}
	response, err := s.AcceptCreditApplication(true, credit)
	if err != nil {
		return model.ResponseCreditApplication{}, err
	}
	errSaveCredit := s.domain.CreditApplication(credit)
	if errSaveCredit != nil {
		return model.ResponseCreditApplication{}, errSaveCredit
	}
	return response, nil
}

// GetLoans - retrieves all the loans [date from, date to]
func (s *Service) GetLoans(dateFrom string, dateTo string) ([]model.CreditApplication, error) {
	loans, err := s.domain.GetLoans(dateFrom, dateTo)
	if err != nil {
		return []model.CreditApplication{}, err
	}
	return loans, nil
}

// PostPayment - adds a new register of one pay realized
func (s *Service) PostPayment(ctx context.Context, payment model.RegisterPaymentMade) (model.DebtPayment, error) {
	userLoan, err := s.domain.GetDebt(payment.UserId)
	if err != nil {
		return model.DebtPayment{}, fmt.Errorf(err.Error())
	}

	// validamos que el monto sea inferior o superior al monto de la cuota
	var debt int
	if payment.Amount > userLoan.Ammount {
		return model.DebtPayment{}, fmt.Errorf("El valor a pagar no puede superar el valor del credito")
	}
	if float32(payment.Amount) <= userLoan.Installment || float32(payment.Amount) >= userLoan.Installment {
		debt = userLoan.Ammount - payment.Amount
	}

	debtPayment := model.DebtPayment{}
	debtPayment.LoanId = int(userLoan.ID)
	debtPayment.Debt = debt
	umm := s.domain.SavePayment(ctx, debtPayment)
	if umm != nil {
		return model.DebtPayment{}, umm
	}
	if result := s.DB.Save(&debtPayment); result.Error != nil {
		return model.DebtPayment{}, nil
	}
	// actualizar deuda
	newDebt := model.UserLoans{
		UserId:  userLoan.UserId,
		Ammount: debt,
	}
	if result := s.DB.Model(&userLoan).Updates(newDebt); result.Error != nil {
		return model.DebtPayment{}, result.Error
	}
	return debtPayment, nil
}

// GetDebt -
func (s *Service) GetDebt(date string, target string) (model.Balnace, error) {
	// hasta una fecha dada
	var userLoan []model.UserLoans
	if result := s.DB.Where("target = ?", target).Or("date = ?", date).Find(&userLoan); result.Error != nil {
		return model.Balnace{}, result.Error
	}
	balance := model.Balnace{}
	for _, v := range userLoan {
		balance.Balance += v.Ammount
	}
	fmt.Println(userLoan)
	return balance, nil
}
