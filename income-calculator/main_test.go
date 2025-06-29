package main

import "testing"

func TestProfitCalculation(t *testing.T) {
	req := IncomeRequest{Revenue: 10000, Costs: 7000}
	expected := 3000.0
	result := req.Revenue - req.Costs
	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}
