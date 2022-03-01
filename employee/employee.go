package employee

const (
	MaxChildrenCount = 2
)

type Employee struct {
	BasicSalary int      `json:"basic_salary"`
	Children    Children `json:"children"`
}

type Children struct {
	Count int     `json:"count"`
	Data  []Child `json:"data"`
}

type Child struct {
	Age       int `json:"age"`
	Allowance int `json:"allowance"`
}

func GetAllowancePercentage(age int) float64 {
	for _, allowance := range ChildrenAllowances {
		if age >= allowance.AgeFrom && age <= allowance.AgeTo {
			return allowance.AllowancePercentage
		}
	}
	return 0
}

func (s *Employee) SetAllowance() {
	if s.Children.Count == 1 {
		s.Children.Data[0].Allowance = int(MaxAllowancePercentage() * float64(s.BasicSalary))
	}

	for i := len(s.Children.Data) - 1; i >= 0; i-- {
		if i < len(s.Children.Data)-MaxChildrenCount {
			continue
		}

		s.Children.Data[i].Allowance = int(GetAllowancePercentage(s.Children.Data[i].Age) * float64(s.BasicSalary))
	}
}
