package usecases

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
	"github.com/jinzhu/gorm"
)

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetHistoryCredit(UserId uint) (model.UserCredit, error) {
	var userCredit model.UserCredit
	if result := s.DB.First(&userCredit, UserId); result.Error != nil {
		return model.UserCredit{}, result.Error
	}
	return userCredit, nil
}

func (s *Service) HistoryCredit(payload model.CreditApplication) (model.UserCredit, error) {

	var userCredit model.UserCredit
	result := s.DB.First(&userCredit, payload.UserId).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		if payload.Amount > 100000 && payload.Amount <= 500000 {
			if result := s.DB.Save(&model.UserCredit{UserId: payload.UserId, AmountTotal: payload.Amount, Cant: 1}); result.Error != nil {
				return model.UserCredit{}, result.Error
			}
			return userCredit, nil
		}
		return model.UserCredit{}, fmt.Errorf("No puedes solicitar un credito mayor a $500.000")
	}
	// UPDATE
	user, _ := s.GetHistoryCredit(payload.UserId)
	newHistory := model.UserCredit{}
	newHistory.UserId = payload.UserId
	newHistory.AmountTotal = user.AmountTotal + payload.Amount
	newHistory.Cant = user.Cant + 1
	if result := s.DB.Model(&user).Updates(newHistory); result.Error != nil {
		return model.UserCredit{}, result.Error
	}
	return newHistory, nil
}

func (s *Service) AcceptCreditApplication(accept bool, creditApplication model.CreditApplication) (model.ResponseCreditApplication, error) {
	if !accept {
		return model.ResponseCreditApplication{}, fmt.Errorf("El credito no fue aprobado")
	}
	cuota := Installment(creditApplication)
	proceess := model.UserLoans{Date: time.Now().Format(time.RFC3339), UserId: creditApplication.UserId, Ammount: creditApplication.Amount, Installment: cuota, Target: creditApplication.Target}
	loanID, err := s.domain.UserLoans(proceess)
	if err != nil {

	}
	return model.ResponseCreditApplication{LoanId: loanID, Installment: cuota}, nil
}

func (s *Service) FecadeTargetClient(historyCredit model.UserCredit, creditApplication model.CreditApplication) (model.CreditApplication, error) {
	if historyCredit.Cant < 2 && creditApplication.Amount <= 500000 {
		ret, err := NewClient(historyCredit, creditApplication)
		if err != nil {
			return model.CreditApplication{}, err
		}
		return ret, nil
	}
	if historyCredit.Cant >= 2 && historyCredit.Cant < 5 && historyCredit.AmountTotal >= 100000 && historyCredit.AmountTotal <= 3000000 {
		credit, err := FrequentClient(historyCredit, creditApplication)
		if err != nil {
			return model.CreditApplication{}, err
		}
		return credit, nil
	}
	if historyCredit.Cant >= 5 && historyCredit.AmountTotal >= 500000 {
		ret, _ := PremiumClient(historyCredit, creditApplication)
		return ret, nil
	}
	return model.CreditApplication{}, nil
}

func Installment(payload model.CreditApplication) float32 {
	var cuotaMensual float32
	r := float32(0.05 / 12)
	umm := float32(math.Pow((1+float64(r)), payload.Term) - 1)
	cuotaMensual = r + r/umm
	cuaota := cuotaMensual * float32(payload.Amount)
	cuotaFinal := fmt.Sprintf("%.2f", cuaota)
	cu, _ := strconv.ParseFloat(cuotaFinal, 32)
	cuu := float32(cu)
	return cuu
}

func NewClient(history model.UserCredit, payload model.CreditApplication) (model.CreditApplication, error) {

	if payload.Amount < 100000 || payload.Amount > 500000 {
		return model.CreditApplication{}, fmt.Errorf("No puedes solicitar un credito menor a $100.000 o mayor a $500.000")
	}
	rate, _ := strconv.ParseFloat(os.Getenv("NEW_RATE"), 8)
	payload.Rate = float32(rate)
	payload.Target = "NEW"
	payload.Date = time.Now().Format(time.RFC3339)
	return payload, nil
}

func FrequentClient(history model.UserCredit, payload model.CreditApplication) (model.CreditApplication, error) {

	if payload.Amount < 100000 || payload.Amount > 1000000 {
		return payload, fmt.Errorf("No puedes solicitar un credito menor a $100.000 o mayor a $1.000.000")
	}
	FREQUENT_RATE, _ := strconv.ParseFloat(os.Getenv("FREQUENT_RATE"), 32)
	payload.Rate = float32(FREQUENT_RATE)
	payload.Target = "FREQUENT"
	payload.Date = time.Now().Format(time.RFC3339)
	return payload, nil
}

func PremiumClient(history model.UserCredit, payload model.CreditApplication) (model.CreditApplication, error) {

	if payload.Amount < 100000 || payload.Amount > 5000000 {
		return payload, fmt.Errorf("No puedes solicitar un credito menor a $100.000 o mayor $5.000.000")
	}
	PREMIUM_RATE, _ := strconv.ParseFloat(os.Getenv("PREMIUM_RATE"), 8)
	payload.Rate = float32(PREMIUM_RATE)
	payload.Target = "PREMIUM"
	payload.Date = time.Now().Format(time.RFC3339)
	return payload, nil
}
