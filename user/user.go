package user

//CONSTANT
//ENUM
//CUSTOM TYPE
//VALIDATION
type Address string

type User struct {
	FirstName       string
	LastName        string
	Role            string //-> ADMIN, NORMAL
	Gender          string //-> MALE, FEMALE, OTHERS
	EmailID         string
	Phone           int
	HomeAddress     Address
	DeliveryAddress Address
}
