package model

import (
	"github.com/jinzhu/gorm"
)

type CreditApplication struct {
	ID     int `gorm:"primary_key"`
	UserId uint
	Amount int
	Term   float64
	Rate   float32
	Target string
	Date   string
}

type UserCredit struct {
	UserId      uint `gorm:"primary_key"`
	AmountTotal int
	Cant        int
}

type ResponseCreditApplication struct {
	LoanId      uint
	Installment float32
}

type UserLoans struct {
	ID          uint `gorm:"primary_key"`
	Date        string
	UserId      uint
	Ammount     int
	Installment float32
	Target      string
}

type RegisterPaymentMade struct {
	UserId uint
	Amount int
}
type DebtPayment struct {
	gorm.Model
	LoanId int
	Debt   int
}

type Balnace struct {
	Balance int
}
