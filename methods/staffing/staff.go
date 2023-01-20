package staffing

var OverPaidLimit = 20000
var UnderPaidLimit = 30000

type Employees struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employees
}

func (e *Office) All() []Employees {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employees {
	var overpaid []Employees

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}

	return overpaid
}

func (e *Office) Underpaid() []Employees {
	var underpaid []Employees

	for _, x := range e.AllStaff {
		if x.Salary < UnderPaidLimit {
			underpaid = append(underpaid, x)
		}
	}

	return underpaid
}
