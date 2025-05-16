package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/amadejkastelic/mercator-api/pkg/mercatorsi"
)

var (
	listCategoriesPtr = flag.Bool("list-categories", false, "List product categories")
	queryPtr          = flag.String("query", "*", "Search query")
	categoryPtr       = flag.String("category", "", "Category ID to filter by")
	pageSizePtr       = flag.Int("page-size", 10, "Number of results per page")
	pagePtr           = flag.Int("page", 0, "Page number")
	userAgentPtr      = flag.String(
		"user-agent",
		"Mozilla/5.0 (X11; Linux i686; rv:128.5) Gecko/20100101 Firefox/128.5",
		"User agent string",
	)
	timeoutPtr = flag.Int("timeout", 15, "Timeout in seconds")
	sortPtr    = flag.String("sort", "current_price", "Sort field (e.g., _score, rating, current_price, weight, first_time_online)")
	sortDirPtr = flag.String("sort-dir", "asc", "Sort direction (asc or desc)")
)

func main() {
	flag.Parse()

	// Create a new client with custom timeout
	client := mercatorsi.NewClient(
		mercatorsi.WithTimeout(time.Duration(*timeoutPtr)*time.Second),
		mercatorsi.WithUserAgent(*userAgentPtr),
	)

	if *listCategoriesPtr {
		listCategories(client)
	}

	// Perform a search
	resp, err := client.Search(mercatorsi.SearchRequest{
		Filter:     *queryPtr,
		CategoryID: *categoryPtr,
		Limit:      *pageSizePtr,
		From:       *pagePtr,
		Sort: &mercatorsi.Sort{
			Field:     mercatorsi.SortField(*sortPtr),
			Direction: mercatorsi.SortDirection(*sortDirPtr),
		},
	})
	if err != nil {
		log.Fatalf("Error searching: %v", err)
	}

	fmt.Printf("Found %d products\n", len(resp.Products))

	for i, product := range resp.Products {
		fmt.Printf("%d. %s - %s€\n", i+1, product.Data.Name, product.Data.CurrentPrice)
	}
}

func listCategories(client mercatorsi.Client) {
	categories, err := client.Categories()
	if err != nil {
		log.Fatalf("Error fetching categories: %v", err)
	}

	fmt.Println("Product Categories:")
	for i, category := range categories.Values {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, category.Category.Name, category.Category.ID)
	}
}
