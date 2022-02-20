package main

import (
	"fmt"
	"log"

	"github.com/yanoandri/simple-goorm/config"
	"github.com/yanoandri/simple-goorm/models"
	"github.com/yanoandri/simple-goorm/repository"
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

	repo := repository.Repository{Database: db}

	// create a payment
	payment := models.Payment{
		PaymentCode: "XXX-1",
		Name:        "Payment for item #1",
		Status:      "PENDING",
	}

	result, err := repo.CreatePayment(payment)
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Payment created", result)

	// select a payment
	var id string
	fmt.Println("Input payment id : ")
	fmt.Scanln(&id)

	payment, _ = repo.SelectPaymentWIthId(id)
	fmt.Println("Your payment is", payment)

	// update a payment with previous id
	updatedPayment, _ := repo.UpdatePayment(id, models.Payment{
		Status: "PAID",
	})
	fmt.Println("Your payment status now is ", updatedPayment)

	// delete a payment with previous id
	repo.DeletePayment(id)
	fmt.Println("Your payment now is deleted")
}
