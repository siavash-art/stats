package stats

import (
	"github.com/siavash-art/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
		{
			ID: 1000,
			Amount: 100,
			Status: "OK",
		},
		{
			ID: 1001,
			Amount: 300,
			Status: "OK",
		},
		{
			ID: 1002,
			Amount: 520,
			Status: "INPROGRESS",
		},
		{
			ID: 1003,
			Amount: 520,
			Status: "FAIL",
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
			Status: "OK",
		},
		{
			ID: 1001,
			Amount: 250,
			Category: "Agro",
			Status: "INPROGRESS",
		},
		{
			ID: 1002,
			Amount: 520,
			Category: "Bussines",
			Status: "FAIL",			
		},
		{
			ID: 1002,
			Amount: 520,
			Category: "Agro",
			Status: "FAIL",			
		},		
	}
	total := TotalInCategory(payments, "Agro")
 	fmt.Println(total)	
	//Output: 350
}	

