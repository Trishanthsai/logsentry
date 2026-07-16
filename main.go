package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	loadEnv()
	var totalrequest int
	var error404 int
	var error5xx int
	var stats []IPStat
	var endpointStats []EndpointStat
	var statusStats []StatusStat
	var severity string
	threshold := 100
	file, err := os.Open("test.txt")
	file1,err1:=os.Open("blacklist.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	
	if err1 != nil {
		fmt.Println("Error opening blacklist file:", err1)
		return
	}
	defer file.Close()
	defer file1.Close()
	scanner:= bufio.NewScanner(file)
	scanner1:= bufio.NewScanner(file1)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if err := scanner1.Err(); err != nil {
		fmt.Println("Error reading blacklist file:", err)
		return
	}
	ipcount:=make(map[string]int)
	endpointcount:=make(map[string]int)
	statusCount := make(map[string]int)
	blacklist := make(map[string]bool)
	alerted:=make(map[string]bool)
	for(scanner1.Scan()){
		ip := strings.TrimSpace(scanner1.Text())
		blacklist[ip] = true
	}
	
	for(scanner.Scan()){
		match := pattern.FindStringSubmatch(scanner.Text())
		if len(match)== 4 {
			totalrequest++;
			ipcount[match[1]]++
			if(blacklist[match[1]] && !alerted[match[1]]) {
				location := getCity(match[1])
				fmt.Println("\n==============================")
				fmt.Println("BLACKLIST ALERT")
				fmt.Println("==============================")
				fmt.Printf("IP Address : %s\n", match[1])
				fmt.Printf("Geo Location   : %s\n", location)
				fmt.Println("==============================")
				sendmail("BLACKLIST ALERT", fmt.Sprintf("IP Address: %s\nGeo Location: %s", match[1], location))
				alerted[match[1]] = true
			}
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
			if !blacklist[ip] {
				addToBlacklist(ip)
				blacklist[ip] = true
				fmt.Println("IP automatically added to blacklist:", ip)
			}
			severity = "LOW"
			if count >= 500 {
				severity = "HIGH"
				fmt.Printf("Sending email alert for IP: %s with %d requests\n", ip, count)
				sendmail("Suspicious Activity Alert", fmt.Sprintf("IP Address: %s\nRequests: %d\nSeverity: %s", ip, count, severity))
			} else if count >= 250 {
				severity = "MEDIUM"
			}		
			// country := getCountry(ip)	
			location:=getCity(ip)
			fmt.Println("\n==============================")
			fmt.Println("Suspicious Activity Detected")
			fmt.Println("==============================")
			fmt.Printf("IP Address : %s\n", ip)
			// fmt.Printf("Country    : %s\n", country)
			fmt.Printf("Requests   : %d\n", count)
			fmt.Printf("Location   : %s\n", location)
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
	if totalrequest > 0 {
		error404Rate = float64(error404) / float64(totalrequest) * 100
		error5xxRate = float64(error5xx) / float64(totalrequest) * 100
	}
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
	report:=Report{
		TotalRequests: totalrequest,
		Error404:      error404,
		Error5xx:      error5xx,
		TopIPs:        stats,
		TopEndpoints:  endpointStats,
		StatusCodes:   statusStats,
	}
	saveReport(report)
	fmt.Println("Watching log file...")
	monitorlogs("test.txt")
}