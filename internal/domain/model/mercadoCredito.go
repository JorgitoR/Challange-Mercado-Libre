package model

type MercadoCredito struct {
}

type CreditApplication struct {
	amount string
	term   string
	userId string
}
type ResponseCreditApplication struct {
	LoanId      string
	Installment string
}

//
type ResponseAllLoan struct {
	amount string
	term   int
	rate   int
	userId int
	target string
	date   string
}

type RegisterPaymentMade struct {
	amount float64
}

type ResponseRegisterPaymentMade struct {
	amount float64
}

type RaiseDebt struct {
	Date string
}

//
type ResponseRaiseDebt struct {
	balance string
}
