package adapters

import (
	"github.com/JorgitoR/Challange-Mercado-Libre/internal/domain/model"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&model.CreditApplication{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&model.UserCredit{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&model.DebtPayment{}); result.Error != nil {
		return result.Error
	}
	if result := db.AutoMigrate(&model.UserLoans{}); result.Error != nil {
		return result.Error
	}
	return nil
}
