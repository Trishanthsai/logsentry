package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)
func main(){
	var totalrequest int
	var error404 int
	var error5xx int
	type IpStat struct {
		IP    string
		Count int
	}
	var stats []IpStat
	file, err := os.Open("test.txt")
	pattern :=regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[.*?\]\s+".*?"\s+(\d{3})`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	ipcount:=make(map[string]int)
	for(scanner.Scan()){
		match := pattern.FindStringSubmatch(scanner.Text())
		if len(match) > 2 {
			totalrequest++;
			ipcount[match[1]]++
			if match[2] == "404" {
				error404++;
			} else if strings.HasPrefix(match[2], "5") {
				error5xx++;
			}
		}

	}
	for ip, count := range ipcount {
		stats = append(stats, IpStat{IP: ip, Count: count})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})
	fmt.Println("File opened successfully:", file.Name())
	fmt.Println("Total Requests:", totalrequest)	
	fmt.Println("Total 404 Errors:", error404)
	fmt.Println("Total 5xx Errors:", error5xx)
	fmt.Println("Top IP Addresses:")
	for i := 0; i < len(stats) && i < 5; i++ {
		fmt.Printf("IP: %s, Count: %d\n", stats[i].IP, stats[i].Count)
	}
}