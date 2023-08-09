package main

import (
	"fmt"
	"math/rand"

	"github.com/boomnoob/go-practice-sql/database"
	"github.com/boomnoob/go-practice-sql/model"
)

func main() {
	database.ConnectDatabase()

	// Create 20 initial customer name
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hannah", "Isaac", "Jack", "Kelly", "Liam", "Mia", "Noah", "Olivia", "Peter", "Quinn", "Ryan", "Sophia", "Tyler"}

	customers := make([]model.Customers, 20)
	for i := 0; i < 20; i++ {
		customers[i] = model.Customers{
			Name: names[i],
			Age:  uint(rand.Intn(41) + 18), // Generates a random age between 18 and 60
			//rand.Intn(max - min + 1) + min
		}
	}

	fmt.Println(customers)

	result := database.DB.Create(&customers)
	if result.Error != nil {
		fmt.Println(`Cannot create new customer`)
		return

	}

}
