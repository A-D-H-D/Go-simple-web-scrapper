package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	//importing Colly
	"github.com/gocolly/colly"
)

// initialize a data structure to kep thr scraped data
type Product struct {
	Url, Image, Name, Price string
}

func main() {
	fmt.Println("hello world")

	// instantiate a new collector object
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	//initialisation of the slice of structs that will contain the scraped data
	var products []Product

	// on HTML callback
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initialize a new product instance
		product := Product{}

		// scrape the target data
		product.Url = e.ChildAttr("a", "href")
		product.Image = e.ChildAttr("img", "src")
		product.Name = e.ChildText(".product-name")
		product.Price = e.ChildText(".price")

		//add the product instance with scraped data to the list of products
		products = append(products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		//open the csv file
		file, err := os.Create("products.csv")

		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		// initialize a file Writer
		writer := csv.NewWriter(file)

		// write the CSV headers
		headers := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}
		writer.Write(headers)

		//write each product as a CSV row

		for _, product := range products {
			// converting a product into an array of strings
			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}

			// add a csv record to the putput file
			writer.Write(record)
		}
		defer writer.Flush()
	})

	// opne the target url
	c.Visit("https://www.scrapingcourse.com/ecommerce")
}