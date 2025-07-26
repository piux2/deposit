package main
import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)
import "regexp"

var storageRe = regexp.MustCompile(`storage:\s*(\d+)`)
var depositRe = regexp.MustCompile(`deposit:\s*(\d+)`)

func main() {
	baseURL := "http://127.0.0.1:8888"
	realms, err := extractRLinks(baseURL+"/r/")
	if err != nil {
		fmt.Println("Failed to extract paths:", err)
		return
	}
	printDeposit(realms)
	println()
	packages, err :=extractRLinks(baseURL+"/p/")
	if err != nil {
		fmt.Println("Failed to extract paths:", err)
		return
	}
	printDeposit(packages)
}

func printDeposit(paths []string){
	var sumStorage, sumDeposit int
	var count int
	fmt.Printf("%-45s %10s %10s\n","Path","Storage (byte)","Deposit (ugnot)")
	for _, path := range paths {
		fullPath := "gno.land" + path

		storage, deposit := queryStorage(fullPath)
		if storage == -1 || deposit == -1 {
			continue
		} 
		fmt.Printf("%-45s %10d %10d\n",fullPath, storage, deposit)
		sumStorage += storage
		sumDeposit += deposit
		count++
	}

	avgStorage := sumStorage / count
	avgDeposit := sumDeposit / count

	println()
	fmt.Println("========== Totals ==========")
	fmt.Printf("Paths counted: %d\n", count)
	fmt.Printf("Total storage: %d byte\n", sumStorage)
	fmt.Printf("Total deposit: %d ugnot\n", sumDeposit)
	fmt.Println("========== Averages ==========")
	fmt.Printf("Average storage: %d byte\n", avgStorage)
	fmt.Printf("Average deposit: %d ugnot\n", avgDeposit)
}
func extractRLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	seen := map[string]bool{}
	var links []string

	var walk func(n *html.Node, insideUL bool)
	walk = func(n *html.Node, insideUL bool) {
		// Entering a <ul> block
		if n.Type == html.ElementNode && n.Data == "ul" {
			insideUL = true
		}

		// Inside a <ul>, look for <a> with href="/r/..."
		if insideUL && n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" &&
   (strings.HasPrefix(attr.Val, "/r/") || strings.HasPrefix(attr.Val, "/p/")) {

					href := attr.Val
					if !seen[href] {
						seen[href] = true
						links = append(links, href)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c, insideUL)
		}
	}

	walk(doc, false)
	return links, nil
}
func queryStorage(path string) (int, int) {
	cmd := exec.Command("gnokey", "query", "vm/qstorage", "--data", path)
	output, err := cmd.Output()
	if err != nil {
		return -1, -1
	}
	text := string(output)
	storage := extractNumber(storageRe, text)
	deposit := extractNumber(depositRe, text)
	return storage, deposit
}
func extractNumber(re *regexp.Regexp, text string) int {
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return -1
	}
	n, err := strconv.Atoi(match[1])
	if err != nil {
		return -1
	}
	return n
}
