package order

import "time"

type Address struct {
	HouseNo     int32
	StreetNo    int32
	AreaVillage string
	City        string
	State       string
	Landmark    string
	PinCode     int64
}
type Order struct {
	OrderNumber     int64
	OrderBy         string
	OrderDate       time.Time
	DeliveryAddress Address
	OrderDelivered  bool
}
