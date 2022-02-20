package repository

import (
	"errors"

	"github.com/yanoandri/simple-goorm/models"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	UpdatePayment(id string, payment models.Payment) (models.Payment, error)
	DeletePayment(id string) (int64, error)
	SelectPaymentWIthId(id string) (models.Payment, error)
	CreatePayment(payment models.Payment) (int64, error)
}

type Repository struct {
	Database *gorm.DB
}

func (repo Repository) UpdatePayment(id string, payment models.Payment) (models.Payment, error) {
	var updatePayment models.Payment
	result := repo.Database.Model(&updatePayment).Where("id = ?", id).Updates(payment)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not update")
	}
	return updatePayment, nil
}

func (repo Repository) DeletePayment(id string) (int64, error) {
	var deletedPayment models.Payment
	result := repo.Database.Where("id = ?", id).Delete(&deletedPayment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment data not update")
	}
	return result.RowsAffected, nil
}

func (repo Repository) SelectPaymentWIthId(id string) (models.Payment, error) {
	var payment models.Payment
	result := repo.Database.First(&payment, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not found")
	}
	return payment, nil
}

func (repo Repository) CreatePayment(payment models.Payment) (int64, error) {
	result := repo.Database.Create(&payment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment not created")
	}
	return result.RowsAffected, nil
}
