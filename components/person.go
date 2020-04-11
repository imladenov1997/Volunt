package components

type Person struct {
	Bill *PersonalBill
}

type PersonalBill struct {
	BillCurrency *Currency 
	Value float64
}