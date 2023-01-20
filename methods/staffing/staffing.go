package staffing

import "log"

var employees = []Employees{
	{
		FirstName: "John",
		LastName:  "Smith",
		Salary:    100000,
		FullTime:  true,
	}, {
		FirstName: "Allan",
		LastName:  "Gurnworthy",
		Salary:    20000,
		FullTime:  false,
	}, {
		FirstName: "Becky",
		LastName:  "Shea",
		Salary:    70000,
		FullTime:  false,
	}, {
		FirstName: "Sinker",
		LastName:  "Swimz",
		Salary:    50000,
		FullTime:  true,
	},
}

func Staffing() {
	myStaff := Office{
		AllStaff: employees,
	}

	//Danger of accessibility here...
	OverPaidLimit = 20000
	log.Println(myStaff.All())
	log.Println(myStaff.Overpaid())
	log.Println(myStaff.Underpaid())
}
