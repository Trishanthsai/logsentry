package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)
type IPStat struct {
		IP    string
		Count int
	}
	type EndpointStat struct {
    Endpoint string
    Count    int
}
type StatusStat struct {
    Status string
    Count  int
}
func main(){
	var totalrequest int
	var error404 int
	var error5xx int
	var stats []IPStat
	var endpointStats []EndpointStat
	var statusStats []StatusStat
	var severity string
	threshold := 100
	file, err := os.Open("test.txt")
	pattern :=regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[.*?\]\s+"\S+\s+(\S+)\s+\S+"\s+(\d{3})`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	ipcount:=make(map[string]int)
	endpointcount:=make(map[string]int)
	statusCount := make(map[string]int)
	for(scanner.Scan()){
		match := pattern.FindStringSubmatch(scanner.Text())
		if len(match)== 4 {
			totalrequest++;
			ipcount[match[1]]++
			endpointcount[match[2]]++
			statusCount[match[3]]++
			if match[3] == "404" {
				error404++;
			} else if strings.HasPrefix(match[3], "5") {
				error5xx++;
			}
		}
	}
	for ip, count := range ipcount {
		if count > threshold {
			severity = "LOW"
			if count >= 500 {
				severity = "HIGH"
			} else if count >= 250 {
				severity = "MEDIUM"
			}			
			fmt.Println("\n==============================")
			fmt.Println("⚠ Suspicious Activity Detected")
			fmt.Println("==============================")
			fmt.Printf("IP Address : %s\n", ip)
			fmt.Printf("Requests   : %d\n", count)
			fmt.Printf("Threshold  : %d\n", threshold)
			fmt.Printf("Severity   : %s\n", severity)
			fmt.Println("==============================")
		}
	}
	for ip, count := range ipcount {
		stats = append(stats, IPStat{IP: ip, Count: count})
	}
	for endpoint, count := range endpointcount {
		endpointStats = append(endpointStats, EndpointStat{Endpoint: endpoint, Count: count})
	}
	for status, count := range statusCount {
		statusStats = append(statusStats, StatusStat{Status: status, Count: count})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})
	sort.Slice(endpointStats, func(i, j int) bool {
		return endpointStats[i].Count > endpointStats[j].Count
	})
	sort.Slice(statusStats, func(i, j int) bool {
		return statusStats[i].Count > statusStats[j].Count
	})
	fmt.Println("File opened successfully:", file.Name())
	fmt.Println("Total Requests:", totalrequest)	
	error404Rate := float64(error404) / float64(totalrequest) * 100
	error5xxRate := float64(error5xx) / float64(totalrequest) * 100
	fmt.Printf("404 Errors : %d (%.2f%%)\n", error404, error404Rate)
	fmt.Printf("5xx Errors : %d (%.2f%%)\n", error5xx, error5xxRate)
	fmt.Println("Top IP Addresses:")
	fmt.Println("IP Address\tCount")
	for i := 0; i < len(stats) && i < 5; i++ {
		fmt.Printf("%s\t%d\n", stats[i].IP, stats[i].Count)
	}
	fmt.Println("Top Endpoints:")
	fmt.Println("Endpoint\tCount")
	for i := 0; i < len(endpointStats) && i < 5; i++ {
		fmt.Printf("%s\t%d\n", endpointStats[i].Endpoint, endpointStats[i].Count)
	}
	fmt.Println("\nStatus Code Distribution")
	fmt.Println("------------------------")
	for i := 0; i < len(statusStats) && i < 5; i++ {
		fmt.Printf("%s\t%d\n", statusStats[i].Status, statusStats[i].Count)
	}
}