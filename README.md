# Programming Exercises Collection

This repository contains three different programming exercises covering web development, API testing, and database analysis.

## 1. Expense Calculator Web Application

A web application to calculate and analyze monthly expenses.

### Features
- Add and manage multiple expenses
- Calculate total monthly expenses
- Calculate average daily expenses
- Display top 3 largest expenses
- Modern, responsive UI

### Files
- `index.html` - Contains the complete web application with HTML, CSS, and JavaScript

### Usage
1. Open `index.html` in a web browser
2. Add expenses by entering category and amount
3. Click "Calculate" to see the analysis

## 2. API Testing with Go

A Go application to test and validate data from the Fake Store API.

### Features
- Validates product data from https://fakestoreapi.com/products
- Checks for data defects:
  - Empty product titles
  - Negative prices
  - Ratings exceeding 5
- Detailed error reporting

### Files
- `api_tests.go` - Main Go application for API testing
- `go.mod` - Go module configuration

### Usage
```bash
go run api_tests.go
```

## 3. SQL Sales Analysis

SQL queries for analyzing sales data from an online store.

### Features
- Calculate total sales for specific periods
- Identify top-spending customers
- Calculate average order values
- Monthly sales breakdown
- Customer purchase frequency analysis

### Files
1. `01_init_database.sql` - Database initialization and data population
2. `02_task1_march_sales.sql` - March 2024 sales calculation
3. `03_task2_top_customer.sql` - Top-spending customer analysis
4. `04_task3_avg_order_value.sql` - Average order value calculation
5. `05_monthly_breakdown.sql` - Monthly sales breakdown
6. `06_customer_analysis.sql` - Customer purchase frequency analysis

### Expected Results
- Total sales for March 2024: 27,000
- Top-spending customer: Alice (20,000)
- Average order value: 6,000

## Setup Instructions

### Expense Calculator
1. Open `index.html` in any modern web browser

### API Testing
1. Install Go (if not already installed)
2. Run `go mod init api_tests`
3. Execute `go run api_tests.go`

### SQL Analysis
1. Run `01_init_database.sql` first to create and populate the database
2. Execute any of the other SQL files to perform specific analyses

## Technologies Used
- HTML/CSS/JavaScript
- Go
- SQLite
- GitHub for version control 