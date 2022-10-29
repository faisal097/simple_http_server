package models

type Emplpoyee struct {
	Name string
	Age  int
	Dept string
}

var PLACE_HOLDER_EMPLOYEES = []Emplpoyee{
	{"Faisal", 28, "IT"}, {"Manish", 45, "IT"},
}
