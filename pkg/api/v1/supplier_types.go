package v1

type Supplier struct {
	SupplierID     uint   `json:"supplierID" gorm:"primary_key"`
	Name           string `json:"name"`
	AddressStreet  string `json:"addressStreet"`
	AddressStreet2 string `json:"addressStreet2"`
	AddressCity    string `json:"addressCity"`
	AddressState   string `json:"addressState"`
	AddressZip     string `json:"addressZip"`
	Phone          string `json:"phone"`
	WebsiteURL     string `json:"website_url"`
	ContactName    string `json:"contactName"`
	ContactPhone   string `json:"contactPhone"`
	ContactEmail   string `json:"contactEmail"`
	Terms          string `json:"terms"`
}
