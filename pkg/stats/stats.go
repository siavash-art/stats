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

// FilterByCategory return payments in category
func FilterByCategory (payments []types.Payment, category types.Category) []types.Payment {
	
	var filtered []types.Payment
	
	for	_, payment := range payments {
		if payment.Category == category {
			filtered = append (filtered, payment)
		}
	}

	return filtered
}

// CategoriesTotal return sum of payment in cayegory
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// CategoriesAvg return avg of payment in category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	
	categories := map[types.Category]types.Money{}
	paymentCountIncategory := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount 
		paymentCountIncategory[payment.Category]++		
	}

	for key := range categories {
		categories[key] /= paymentCountIncategory[key]
	}

	return categories
}