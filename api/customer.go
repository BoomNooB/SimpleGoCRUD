package api

import (
	// build-in

	"net/http"

	//3rd party
	"github.com/gin-gonic/gin"

	//local
	"github.com/boomnoob/go-practice-sql/database"
	"github.com/boomnoob/go-practice-sql/model"
)

// POST
// Create Cumtomer
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
				"error": "WTF",
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

// GET
// GET USER INFO BY ID
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

// DELETE
// delete record by record_id
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

func SetupEndpoint() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "API is running at "})
	})

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

	r.Run(":22356")

}
