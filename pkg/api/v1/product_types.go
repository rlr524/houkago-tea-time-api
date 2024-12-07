package v1

type Product struct {
	ProductID uint    `json:"productID" gorm:"primary_key"`
	Name      string  `json:"name"`
	Cost      float64 `json:"cost"`
	Retail    float64 `json:"retail"`
}
