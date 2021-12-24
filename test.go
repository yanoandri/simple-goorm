package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/yanoandri/simple-goorm/config"
	"github.com/yanoandri/simple-goorm/models"
	"gorm.io/gorm"
)

func main() {
	//connect to postgresql
	db, err := config.Setup()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Connected")
	// migrate models inside project
	db.AutoMigrate(models.Payment{})
	fmt.Println("Migrated")

	// create a payment
	payment := models.Payment{
		PaymentCode: "XXX-1",
		Name:        "Payment for item #1",
		Status:      "PENDING",
	}

	result, err := CreatePayment(db, payment)
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Payment created", result)

	// select a payment
	var id string
	fmt.Println("Input payment id : ")
	fmt.Scanln(&id)

	payment, _ = SelectPaymentWIthId(db, id)
	fmt.Println("Your payment is", payment)

	// update a payment with previous id
	updatedPayment, _ := UpdatePayment(db, id, models.Payment{
		Status: "PAID",
	})
	fmt.Println("Your payment status now is ", updatedPayment)

	// delete a payment with previous id
	DeletePayment(db, id)
	fmt.Println("Your payment now is deleted")
}

func UpdatePayment(db *gorm.DB, id string, payment models.Payment) (models.Payment, error) {
	var updatePayment models.Payment
	result := db.Model(&updatePayment).Where("id = ?", id).Updates(payment)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not update")
	}
	return updatePayment, nil
}

func DeletePayment(db *gorm.DB, id string) (int64, error) {
	var deletedPayment models.Payment
	result := db.Where("id = ?", id).Delete(&deletedPayment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment data not update")
	}
	return result.RowsAffected, nil
}

func SelectPaymentWIthId(db *gorm.DB, id string) (models.Payment, error) {
	var payment models.Payment
	result := db.First(&payment, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not found")
	}
	return payment, nil
}

func CreatePayment(db *gorm.DB, payment models.Payment) (int64, error) {
	result := db.Create(&payment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment not created")
	}
	return result.RowsAffected, nil
}
