package stats

import "github.com/siavash-art/bank/v2/pkg/types"

//Avg get average payment
func Avg(payments []types.Payment) types.Money {
	
	total := types.Money(0)
	pCount := types.Money(0)
	for _, payment := range payments {		
		
		if payment.Status != types.StatusFail {
			total += payment.Amount		
			pCount++
		}					
	}

	average := total / pCount

	return average
}

//TotalInCategory total payment of category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	total := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			total += payment.Amount
		}
	}

	return total
}

