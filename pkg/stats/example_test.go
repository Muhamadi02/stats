package stats

import (
	"fmt"
	"github.com/Muhamadi02/banktypes/v2/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   2000_00,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   3000_00,
			Category: "medicine",
			Status:   types.StatusFail,
		},
		{
			ID:       3,
			Amount:   1000_00,
			Category: "restaurants",
			Status:   types.StatusOk,
		},
	}

	fmt.Println(Avg(payments))
	// Output:
	// 150000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   2000_00,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   3000_00,
			Category: "medicine",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   1000_00,
			Category: "restaurants",
			Status:   types.StatusFail,
		},
		{
			ID:       4,
			Amount:   5000_00,
			Category: "auto",
			Status:   types.StatusOk,
		},
	}

	fmt.Println(TotalInCategory(payments, "auto"))
	fmt.Println(TotalInCategory(payments, "medicine"))
	fmt.Println(TotalInCategory(payments, "restaurants"))

	// Output:
	// 700000
	// 300000
	// 0
}
