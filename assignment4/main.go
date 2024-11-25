package main

import "fmt"

type ExpenseTracker struct {
	CategoryExpense map[string]float64
	TotalExpense    float64
	recentExpenses  []float64
}

func initialise() *ExpenseTracker {
	return &ExpenseTracker{
		CategoryExpense: make(map[string]float64),
		recentExpenses:  make([]float64, 3),
	}
}

func (e *ExpenseTracker) Totalexpense() func(amount float64) float64 {
	return func(amount float64) float64 {
		e.TotalExpense += amount
		return e.TotalExpense
	}

}

func (e *ExpenseTracker) addExpense(expense string, amount float64) float64 {
	e.CategoryExpense[expense] += amount
	e.recentExpenses = e.recentExpenses[1:]
	e.recentExpenses = append(e.recentExpenses, amount)
	totalexp := e.Totalexpense()
	return totalexp(amount)
}
func (e *ExpenseTracker) viewexpenses() {
	for expense, price := range e.CategoryExpense {
		fmt.Printf("Expense in %s is $%.2f\n", expense, price)
	}
}
func (e *ExpenseTracker) viewrecentexpenses() {
	fmt.Println(e.recentExpenses)
}
func main() {
	tracker := initialise()
	total := tracker.Totalexpense()
	fmt.Printf("%T", total)
	//tracker.addExpense("Gas", 100.0)

}
