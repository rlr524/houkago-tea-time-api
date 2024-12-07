package v1

type Category struct {
	CategoryID uint   `json:"categoryID" gorm:"primary_key"`
	Name       string `json:"name"`
}
