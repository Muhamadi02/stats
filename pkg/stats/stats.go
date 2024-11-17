package stats

import "github.com/Muhamadi02/banktypes/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	t := 0
	for _, payment := range payments {
		sum += payment.Amount
		t++
	}

	return sum / types.Money(t)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
