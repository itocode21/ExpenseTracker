package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// --------------------------------------------------
// Expense struct
type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"text"`
	Amount      int       `json:"amount"`
	ExpenseTime time.Time `json:"expense_time"`
}

var expenses []Expense

//--------------------------------------------------
// read .json file

func loadExpense() {
	file, err := os.ReadFile("expense.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, &expenses)
}

// --------------------------------------------------
//savew new expense

func saveExpense() {
	file, _ := json.MarshalIndent(expenses, "", "	")
	_ = os.WriteFile("expense.json", file, 0644)
}

//--------------------------------------------------
// add expense

func addExpense(text string, amount int) {
	expense := Expense{
		ID:          len(expenses) + 1,
		Description: text,
		Amount:      amount,
		ExpenseTime: time.Now()}

	expenses = append(expenses, expense)
	saveExpense()
	fmt.Printf("Expense added successfully [ID %d]", expense.ID)
}

//--------------------------------------------------
// delete expense

func deleteExpense(id int) {
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			saveExpense()
			return
		}
	}
}

// --------------------------------------------------
// list all expense

func listExpense() {
	fmt.Println("# ID  Date       Description  Amount")
	for _, expense := range expenses {
		fmt.Printf("# %d   %s  %s       $%d\n",
			expense.ID, expense.ExpenseTime.Format("2006-01-02"),
			expense.Description,
			expense.Amount)
	}
}

func main() {
	loadExpense()
	if len(os.Args) < 2 {
		fmt.Println("No Command provided")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: add [description] [amount]")
			return
		}
		text := os.Args[2]
		amount, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid amount. It must be a number.")
			return
		}
		addExpense(text, amount)
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteExpense(id)
	case "list":
		listExpense()
	default:
		fmt.Println("Invalid command")
	}
}
