package employee

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAllowance(t *testing.T) {

	t.Run("case #1", func(t *testing.T) {
		employee := &Employee{
			Children: Children{
				Count: 2,
				Data:  []Child{{Age: 10}, {Age: 15}},
			},
			BasicSalary: 3500000,
		}
		employee.SetAllowance()

		b, _ := json.Marshal(employee)
		fmt.Println(string(b))

		allowanceTotal := 0
		for _, child := range employee.Children.Data {
			allowanceTotal += child.Allowance
		}

		require.Equal(t, 595000, allowanceTotal)
	})

	t.Run("case #2", func(t *testing.T) {
		employee := &Employee{
			Children: Children{
				Count: 0,
				Data:  []Child{},
			},
			BasicSalary: 2700000,
		}
		employee.SetAllowance()

		b, _ := json.Marshal(employee)
		fmt.Println(string(b))

		allowanceTotal := 0
		for _, child := range employee.Children.Data {
			allowanceTotal += child.Allowance
		}

		require.Equal(t, 0, allowanceTotal)
	})

	t.Run("case #3", func(t *testing.T) {
		employee := &Employee{
			Children: Children{
				Count: 4,
				Data:  []Child{{Age: 3}, {Age: 5}, {Age: 13}, {Age: 14}},
			},
			BasicSalary: 5000000,
		}
		employee.SetAllowance()

		b, _ := json.Marshal(employee)
		fmt.Println(string(b))

		allowanceTotal := 0
		for _, child := range employee.Children.Data {
			allowanceTotal += child.Allowance
		}

		require.Equal(t, 1000000, allowanceTotal)
	})

	t.Run("case #4", func(t *testing.T) {
		employee := &Employee{
			Children: Children{
				Count: 1,
				Data:  []Child{{Age: 6}},
			},
			BasicSalary: 4000000,
		}
		employee.SetAllowance()

		b, _ := json.Marshal(employee)
		fmt.Println(string(b))

		allowanceTotal := 0
		for _, child := range employee.Children.Data {
			allowanceTotal += child.Allowance
		}

		require.Equal(t, 280000, allowanceTotal)
	})

}
