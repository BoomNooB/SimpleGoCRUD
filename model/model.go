package model

type Customers struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
