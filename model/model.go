package model

type Customers struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
