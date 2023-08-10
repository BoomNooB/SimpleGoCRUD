package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boomnoob/go-practice-sql/database"
	"github.com/boomnoob/go-practice-sql/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ReadinessCheck(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": "API is running"})
}

func GetCustomerInfo(c *gin.Context) {
	var customer model.Customers
	err := database.DB.Where("id = ?", c.Param("id")).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Cannot find customer information",
			})
		return
	}

	c.JSON(http.StatusOK, customer)

}

func DeleteCustomer(c *gin.Context) {
	var customer model.Customers

	err := database.DB.Where("id = ?", c.Param("id")).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Cannot find customer information",
			})
		return
	}

	database.DB.Delete(&customer)

	c.JSON(http.StatusOK, customer)

}

func CreateNewCustomer(c *gin.Context) {
	var customer model.Customers

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Invalid customer data",
			})
		return
	}

	if customer.Name == "" || customer.Age == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Name and age must be fill, and age must be only number",
			})
		return
	}

	result := database.DB.Create(&customer)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Cannot create new customer",
			})
		return

	}

	c.JSON(http.StatusCreated, &customer)
}

// PUT
// UPDATE EVERY FIELD BY ID
// if ID not found then return 404

func UpdateCustomerInfo(c *gin.Context) {
	var customer model.Customers

	err := database.DB.Where("id = ?", c.Param("id")).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Cannot find customer information",
			})
		return
	}

	err = c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Invalid customer data",
			})
		return
	}

	database.DB.Model(&customer).Updates(&customer)

	c.JSON(http.StatusOK, customer)

}

func setupRouter() *gin.Engine {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(".env loaded")

	r := gin.Default()

	r.GET("/", ReadinessCheck)

	customersGroup := r.Group("/customers")
	{
		customersGroup.POST("", CreateNewCustomer)
		customersGroup.PUT("/:id", UpdateCustomerInfo)
	}

	employeeGroup := r.Group("/employees")
	{
		employeeGroup.GET("/:id", GetCustomerInfo)
		employeeGroup.DELETE("/:id", DeleteCustomer)
	}

	database.ConnectDatabase()

	return r
}

func Main() {
	r := setupRouter()
	r.Run(":" + os.Getenv("API_PORT"))
}
