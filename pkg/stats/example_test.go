package stats

import (
	"github.com/siavash-art/bank/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
		{
			ID: 1000,
			Amount: 100,
		},
		{
			ID: 1001,
			Amount: 300,
		},
		{
			ID: 1002,
			Amount: 520,
		},		
	}

 	fmt.Println(Avg(payments))	
	//Output: 306
}

func ExampleTotalInCategory() {
	payments := []types.Payment {
		{
			ID: 1000,
			Amount: 100,
			Category: "Agro",
		},
		{
			ID: 1001,
			Amount: 250,
			Category: "Agro",
		},
		{
			ID: 1002,
			Amount: 520,
			Category: "Bussines",
		},		
	}
	total := TotalInCategory(payments, "Agro")
 	fmt.Println(total)	
	//Output: 350
}	

