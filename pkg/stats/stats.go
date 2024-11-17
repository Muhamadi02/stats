package stats

import "github.com/Muhamadi02/banktypes/v2/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	t := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			t++
		}
	}

	return sum / types.Money(t)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return sum
}
