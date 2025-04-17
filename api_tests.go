package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Product represents the structure of a product from the API
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

// Rating represents the rating structure within a product
type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

// Defect represents a validation issue found in a product
type Defect struct {
	ProductID int
	Title     string
	Issues    []string
}

const apiURL = "https://fakestoreapi.com/products"

func main() {
	fmt.Println("Starting API tests...\n")

	// Test 1: Verify server response code
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	fmt.Println("Test 1: Server Response Code")
	fmt.Printf("Expected: 200, Actual: %d\n", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		fmt.Println("✅ Passed")
	} else {
		fmt.Println("❌ Failed")
		os.Exit(1)
	}
	fmt.Println("------------------------\n")

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	// Parse JSON response
	var products []Product
	if err := json.Unmarshal(body, &products); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Test 2: Validate product data
	fmt.Println("Test 2: Product Data Validation")
	defects := validateProducts(products)

	if len(defects) == 0 {
		fmt.Println("✅ No defects found in any products")
	} else {
		fmt.Printf("❌ Found %d products with defects:\n", len(defects))
		for _, defect := range defects {
			fmt.Printf("\nProduct ID: %d\n", defect.ProductID)
			fmt.Printf("Title: %s\n", defect.Title)
			fmt.Println("Defects:")
			for _, issue := range defect.Issues {
				fmt.Printf("- %s\n", issue)
			}
		}
	}
}

func validateProducts(products []Product) []Defect {
	var defects []Defect

	for _, product := range products {
		var issues []string

		// Check if title is empty
		if product.Title == "" {
			issues = append(issues, "Empty title")
		}

		// Check if price is negative
		if product.Price < 0 {
			issues = append(issues, "Negative price")
		}

		// Check if rating.rate exceeds 5
		if product.Rating.Rate > 5 {
			issues = append(issues, "Rating exceeds 5")
		}

		if len(issues) > 0 {
			defects = append(defects, Defect{
				ProductID: product.ID,
				Title:     product.Title,
				Issues:    issues,
			})
		}
	}

	return defects
} 