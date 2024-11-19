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

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// CategoriesAvg возвращает среднюю сумму по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	countCategories := map[types.Category]int{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		countCategories[payment.Category]++
	}

	for key := range categories {
		categories[key] /= types.Money(countCategories[key])
	}

	return categories
}
