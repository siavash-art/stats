package stats

import (
	"reflect"
	"testing"
	"github.com/siavash-art/bank/v2/pkg/types"
)

func TestFilterByCategory_nil(t *testing.T) {

	var payments []types.Payment

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}	
}

func TestFilterByCategory_empty(t *testing.T) {

	payments := []types.Payment {}

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
	
}

func TestFilterByCategory_notFound(t *testing.T) {

	payments := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
	
}

func TestFilterByCategory_foundMultiple(t *testing.T) {

	payments := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	
	expected := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}	
}

func TestFilterByCategories_Total(t *testing.T) {

	payments := []types.Payment {
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	
	expected := map[types.Category]types.Money {
		"auto": 8_000_00,
		"food": 2_000_00,
		"fun": 5_000_00,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}	
}

func TestFilterByCategories_Avg(t *testing.T) {

	payments := []types.Payment {
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	
	expected := map[types.Category]types.Money {
		"auto": 2_666_66,
		"food": 2_000_00,
		"fun": 5_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}	
}

func TestFilterByCategories_PeriodsDynamic(t *testing.T) {

	firstPayments := map[types.Category]types.Money {
		"auto": 10,
		"food": 20,				
	}
	secondPayments := map[types.Category]types.Money {
		"auto": 10,
		"food": 25,			
		"mobile": 5,			
	}
	expected := map[types.Category]types.Money {
		"auto": 0,
		"food": 5,
		"mobile": 5,		
	}

	result := PeriodsDynamic(firstPayments, secondPayments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v, actual: %v", expected, result)
	}	
}