package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := getInput("Revenue: ")
	expenses, err2 := getInput("Expenses: ")
	taxRate, err3 := getInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Print(err1)
		return
	}

	earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeValues(earningsBeforeTax, profit, ratio)

}

func getInput(infoText string) (float64, error) {
	var value float64
	fmt.Print(infoText)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("os valores não podem ser menores que 1")
	}
	return value, nil
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func storeValues(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
