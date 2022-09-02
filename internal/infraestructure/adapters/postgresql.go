package adapters

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DTBAdapter struct {
	DB  *gorm.DB
	DBP *sql.DB
}

func NewPostgreSQLAdapter(dd *gorm.DB, db *sql.DB) *DTBAdapter {
	return &DTBAdapter{
		DB:  dd,
		DBP: db,
	}
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

func (s *DTBAdapter) SavePaymentDatabase(ctx context.Context, args model.DebtPayment) error {
	return nil
}

///////////////////
func (s *DTBAdapter) CreditApplication(creditApplication model.CreditApplication) error {
	fmt.Println(time.Now().Format(time.RFC3339))
	_, credit := s.FindCreditApplication(creditApplication)
	fmt.Println(credit)
	if !credit {
		err := s.SaveCreditApplication(creditApplication)
		return err
	}
	err := s.UpdateCreditApplication(creditApplication)
	return err
}
func (s *DTBAdapter) SaveCreditApplication(creditApplication model.CreditApplication) error {
	if result := s.DB.Save(&creditApplication); result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *DTBAdapter) FindCreditApplication(creditApplication model.CreditApplication) (model.CreditApplication, bool) {
	result := s.DB.First(&creditApplication, "user_id = ?", creditApplication.UserId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.CreditApplication{}, false
	}
	return creditApplication, true
}

func (s *DTBAdapter) UpdateCreditApplication(creditApplication model.CreditApplication) error {
	creditUpdate, _ := s.FindCreditApplication(creditApplication)
	if result := s.DB.Model(&creditUpdate).Updates(creditApplication); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *DTBAdapter) UserLoans(userLoans model.UserLoans) (uint, error) {
	_, loan := s.FindUserLoans(userLoans)
	if !loan {
		loanId, err := s.SaveUserLoan(userLoans)
		return loanId.ID, err
	}
	loanId, err := s.UpdateUserLoan(userLoans)
	return loanId.ID, err
}
func (s *DTBAdapter) FindUserLoans(userLoans model.UserLoans) (model.UserLoans, bool) {
	result := s.DB.First(&userLoans, "user_id = ?", userLoans.UserId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.UserLoans{}, false
	}
	return userLoans, true
}
func (s *DTBAdapter) SaveUserLoan(userLoan model.UserLoans) (model.UserLoans, error) {
	if result := s.DB.Save(&userLoan); result.Error != nil {
		return model.UserLoans{}, result.Error
	}
	return userLoan, nil
}

func (s *DTBAdapter) UpdateUserLoan(newUserLoan model.UserLoans) (model.UserLoans, error) {
	userLoan, _ := s.FindUserLoans(newUserLoan)
	if result := s.DB.Model(&userLoan).Updates(newUserLoan); result.Error != nil {
		return model.UserLoans{}, result.Error
	}
	return newUserLoan, nil
}

func (s *DTBAdapter) GetLoans(dateFrom string, dateTo string) ([]model.CreditApplication, error) {
	var creditApplication []model.CreditApplication
	if result := s.DB.Find(&creditApplication, "date BETWEEN ? AND ?", dateFrom, dateTo); result.Error != nil {
		return []model.CreditApplication{}, result.Error
	}
	return creditApplication, nil
}

func (s *DTBAdapter) RegisterPaymentMade(registerPaymentMade model.RegisterPaymentMade) (model.DebtPayment, error) {
	if result := s.DB.Save(&registerPaymentMade); result.Error != nil {
		return model.DebtPayment{}, result.Error
	}
	return model.DebtPayment{}, nil
}

func (s *DTBAdapter) GetDebt(id uint) (model.UserLoans, error) {
	var userLoans model.UserLoans
	result := s.DB.First(&userLoans, "user_id = ?", id).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		return model.UserLoans{}, result
	}
	return userLoans, nil
}
