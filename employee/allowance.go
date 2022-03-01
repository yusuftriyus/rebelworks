package employee

type Allowance struct {
	AgeFrom             int     `json:"age_from"`
	AgeTo               int     `json:"age_to"`
	AllowancePercentage float64 `json:"allowance_percentage"`
}

var ChildrenAllowances = []Allowance{
	{
		AgeFrom:             1,
		AgeTo:               5,
		AllowancePercentage: 0.05,
	},
	{
		AgeFrom:             6,
		AgeTo:               10,
		AllowancePercentage: 0.07,
	},
	{
		AgeFrom:             11,
		AgeTo:               15,
		AllowancePercentage: 0.1,
	},
}

func MaxAllowancePercentage() float64 {
	max := 0.0
	for _, c := range ChildrenAllowances {
		if c.AllowancePercentage > max {
			max = c.AllowancePercentage
		}
	}
	return max
}
