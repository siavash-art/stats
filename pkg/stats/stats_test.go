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