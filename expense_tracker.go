package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct{
	ID				int  		`json: "id"`
	Description 	string		`json: "descriptio"`
	Amount			float64		`json: "amount"`
	Category		string		`json: "category"`
	Date			time.Time	`json: "date"`
}

type ExpenseTracker struct{
	Expenses []Expense 		`json: "expense"`
	NextID int				`json: "nextid"`
	filename string
}

func NewExpenseTracker(filename string) *ExpenseTracker {
	tracker := &ExpenseTracker{
		Expenses: []Expense{},
		NextID:   1,
		filename: filename,
	}
	tracker.loadFromFile()
	return tracker
}

// loadFromFile loads expenses from JSON file
func (et *ExpenseTracker) loadFromFile() {
	file, err := os.ReadFile(et.filename)
	if err != nil {
		// File doesn't exist, start with empty tracker
		return
	}
	
	json.Unmarshal(file, et)
}

// saveToFile saves expenses to JSON file
func (et *ExpenseTracker) saveToFile() error {
	data, err := json.MarshalIndent(et, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(et.filename, data, 0644)
}

// AddExpense adds a new expense
func (et *ExpenseTracker) AddExpense(description string, amount float64, category string) {
	expense := Expense{
		ID:          et.NextID,
		Description: description,
		Amount:      amount,
		Category:    category,
		Date:        time.Now(),
	}
	
	et.Expenses = append(et.Expenses, expense)
	et.NextID++
	et.saveToFile()
	
	fmt.Printf("✅ Added expense: $%.2f for %s\n", amount, description)
}

// ListExpenses displays all expenses
func (et *ExpenseTracker) ListExpenses() {
	if len(et.Expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}
	
	fmt.Println("\n📊 Your Expenses:")
	fmt.Println("ID | Date       | Category    | Description              | Amount")
	fmt.Println("---|------------|-------------|--------------------------|--------")
	
	total := 0.0
	for _, expense := range et.Expenses {
		fmt.Printf("%-2d | %-10s | %-11s | %-24s | $%.2f\n",
			expense.ID,
			expense.Date.Format("2006-01-02"),
			expense.Category,
			truncateString(expense.Description, 24),
			expense.Amount)
		total += expense.Amount
	}
	
	fmt.Printf("\n💰 Total: $%.2f\n", total)
}

// DeleteExpense removes an expense by ID
func (et *ExpenseTracker) DeleteExpense(id int) {
	for i, expense := range et.Expenses {
		if expense.ID == id {
			et.Expenses = append(et.Expenses[:i], et.Expenses[i+1:]...)
			et.saveToFile()
			fmt.Printf("Deleted expense: %s\n", expense.Description)
			return
		}
	}
	fmt.Printf("Expense with ID %d not found\n", id)
}

// GetExpensesByCategory shows expenses grouped by category
func (et *ExpenseTracker) GetExpensesByCategory() {
	if len(et.Expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}
	
	categoryTotals := make(map[string]float64)
	for _, expense := range et.Expenses {
		categoryTotals[expense.Category] += expense.Amount
	}
	
	fmt.Println("\n Expenses by Category:")
	for category, total := range categoryTotals {
		fmt.Printf("%-15s: $%.2f\n", category, total)
	}
}

// truncateString truncates a string to specified length
func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}

// getUserInput prompts user for input
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// showMenu displays the main menu
func showMenu() {
	fmt.Println("\n Expense Tracker")
	fmt.Println("==================")
	fmt.Println("1. Add expense")
	fmt.Println("2. List all expenses")
	fmt.Println("3. Delete expense")
	fmt.Println("4. View by category")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option (1-5): ")
}

func main() {
	tracker := NewExpenseTracker("expenses.json")
	
	fmt.Println("Welcome to your Personal Expense Tracker!")
	
	for {
		showMenu()
		
		choice := getUserInput("")
		
		switch choice {
		case "1":
			// Add expense
			description := getUserInput("Enter description: ")
			if description == "" {
				fmt.Println("Description cannot be empty")
				continue
			}
			
			amountStr := getUserInput("Enter amount: $")
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Please enter a valid amount")
				continue
			}
			
			category := getUserInput("Enter category (e.g., Food, Transport, Entertainment): ")
			if category == "" {
				category = "Other"
			}
			
			tracker.AddExpense(description, amount, category)
			
		case "2":
			// List expenses
			tracker.ListExpenses()
			
		case "3":
			// Delete expense
			tracker.ListExpenses()
			if len(tracker.Expenses) == 0 {
				continue
			}
			
			idStr := getUserInput("Enter expense ID to delete: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Please enter a valid ID")
				continue
			}
			
			tracker.DeleteExpense(id)
			
		case "4":
			// View by category
			tracker.GetExpensesByCategory()
			
		case "5":
			// Exit
			fmt.Println("Thanks for using Expense Tracker!")
			return
			
		default:
			fmt.Println("Invalid choice. Please select 1-5.")
		}
	}
}