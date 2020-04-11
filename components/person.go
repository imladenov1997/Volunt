package components

type Person struct {
	Bill PersonalBill
}

type PersonalBill struct {
	Currency string
	Value float64
}