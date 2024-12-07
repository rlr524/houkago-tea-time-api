package v1

type Product struct {
	ProductID     uint     `json:"productID" gorm:"primary_key"`
	Name          string   `json:"name"`
	Category      Category `json:"category"`
	Supplier      Supplier `json:"supplier"`
	Cost          float64  `json:"cost"`
	Price         float64  `json:"retail"`
	SalePrice     float64  `json:"salePrice"`
	Size          string   `json:"size"`
	Style         string   `json:"style"`
	Color         string   `json:"color"`
	SupplierModel string   `json:"supplierModel"`
	SupplierID    string   `json:"supplierID"`
}
