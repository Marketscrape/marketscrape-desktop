package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Fetch the HTML content using curl
	cmd := exec.Command("curl",
		"https://www.ebay.ca/sch/i.html?_from=R40&_trksid=p2334524.m570.l1313&_nkw=ipad&_sacat=0&_odkw=asdasdasdas&_osacat=0",
		"--compressed",
		"-H", "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:134.0) Gecko/20100101 Firefox/134.0",
		"-H", "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"-H", "Accept-Language: en-CA,en-US;q=0.7,en;q=0.3",
		"-H", "Accept-Encoding: gzip, deflate, br, zstd",
		"-H", "Referer: https://www.ebay.ca/sch/i.html?_from=R40&_trksid=p4439441.m570.l1313&_nkw=asdasdasdas&_sacat=0",
		"-H", "Connection: keep-alive",
		"-H", "Upgrade-Insecure-Requests: 1",
		"-H", "Sec-Fetch-Dest: document",
		"-H", "Sec-Fetch-Mode: navigate",
		"-H", "Sec-Fetch-Site: same-origin",
		"-H", "Sec-Fetch-User: ?1",
		"-H", "Priority: u=0, i",
		"-H", "TE: trailers",
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", stderr.String())
		return
	}

	// Save HTML to a file for debugging
	htmlFile := "ebay_result.html"
	err = os.WriteFile(htmlFile, out.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(&out)
	if err != nil {
		fmt.Printf("Error parsing HTML: %s\n", err)
		return
	}

	// Open an output file for saving results
	outputFile := "ebay_results_parsed.txt"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %s\n", err)
		return
	}
	defer file.Close()

	// Extract title, price, and link
	doc.Find(".s-item").Each(func(index int, item *goquery.Selection) {
		title := item.Find(".s-item__title").Text()
		link, _ := item.Find(".s-item__link").Attr("href")
		price := item.Find(".s-item__price").Text()

		if title != "" && link != "" && price != "" {
			result := fmt.Sprintf("Title: %s\nLink: %s\nPrice: %s\n\n", title, link, price)
			_, err := file.WriteString(result)
			if err != nil {
				fmt.Printf("Error writing result: %s\n", err)
			}
		}
	})

	fmt.Printf("Scraping complete. Results saved in %s\n", outputFile)
}
