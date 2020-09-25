package stats

import ("github.com/siavash-art/bank/pkg/types")

//Avg get average payment
func Avg(payments []types.Payment) types.Money {

	total := types.Money(0)

	for _, payment := range payments {
		total += payment.Amount
	}

	average := total / types.Money(len(payments))

	return average
}

//TotalInCategory total payment of category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	total := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			total += payment.Amount
		}
	}

	return total
}

