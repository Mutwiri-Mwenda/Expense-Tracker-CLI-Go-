# Expense Tracker CLI (Go)
A simple command-line expense tracker built in Go.
It helps you log, view, and manage your daily expenses with categorized summaries — all stored locally in a JSON file.

# Features
Add new expenses with description, amount, and category

View all expenses with dates and totals

Delete expenses by ID

View expense totals by category

Automatically saves to a local expenses.json file

# How to Run
1. Clone the Repo
bash
git clone https://github.com/Mutwiri-Mwenda/Expense-Tracker-CLI-Go-.git
cd expense-tracker-cli
2. Build and Run
bash
go run main.go
Or build a binary:

bash
go build -o tracker
./tracker
Sample Output
bash
Welcome to your Personal Expense Tracker!

Expense Tracker
==================
1. Add expense
2. List all expenses
3. Delete expense
4. View by category
5. Exit
Choose an option (1-5):
# Data Storage

All data is stored in expenses.json in the current directory. This allows you to:

Resume tracking where you left off

View/edit the raw file manually if needed

# Built With

Go – standard libraries only (no dependencies)

JSON for persistent local storage

CLI-based UI

# Future Improvements

Monthly reports

Edit existing expenses

CSV export

Category-based budgeting


