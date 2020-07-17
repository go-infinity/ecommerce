package user

import "time"

//CONSTANT
const (
	INDIA string = "India"
	UP    string = "Uttar Pradesh"
)

//ENUM
//CUSTOM TYPE
//VALIDATION
type Address struct {
	City    string
	State   string
	Country string
}

type User struct {
	FirstName       string
	LastName        string
	Role            string //-> ADMIN, NORMAL
	Gender          string //-> MALE, FEMALE, OTHERS
	EmailID         string
	Phone           int
	HomeAddress     Address
	DeliveryAddress Address
	IsActive        bool
	LastLoginTime   time.Time
}
