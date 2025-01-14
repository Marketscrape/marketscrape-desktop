package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
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

	outputFile := "ebay_result.html"
	err = os.WriteFile(outputFile, out.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	fmt.Printf("HTML output saved to %s\n", outputFile)
}
