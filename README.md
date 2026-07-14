# LogSentry

A lightweight command-line server log analyzer built with Go that processes web server access logs, extracts traffic insights, detects HTTP errors, ranks the most active client IP addresses, and generates an audit report.

---

## Features

- Read server log files line-by-line
- Extract client IP addresses using Regular Expressions
- Extract HTTP status codes
- Count total requests
- Detect HTTP 404 (Not Found) errors
- Detect HTTP 5xx (Server) errors
- Track request frequency per IP address
- Display the Top 3 most active client IPs
- Fast in-memory processing using Go maps

---

## Tech Stack

- Go
- Regular Expressions (`regexp`)
- File I/O
- Buffered Scanning (`bufio.Scanner`)
- Maps
- Structs
- Slices
- Custom Sorting (`sort.Slice`)

---

## Project Structure

```
logsentry/
│
├── main.go
├── go.mod
├── server.log
├── README.md
└── screenshots/
```

---

## How It Works

```
Server Log
     │
     ▼
Read Line by Line
     │
     ▼
Regex Extraction
     │
     ▼
Extract IP + Status Code
     │
     ▼
Count Requests
     │
     ▼
Track HTTP Errors
     │
     ▼
Aggregate Statistics
     │
     ▼
Sort by Request Count
     │
     ▼
Generate Audit Report
```

---

## Example Input

```
192.168.1.10 - - [14/May/2026:10:00:01 +0000] "GET /index.html HTTP/1.1" 200 1024
10.0.0.5 - - [14/May/2026:10:00:02 +0000] "GET /login HTTP/1.1" 404 512
203.0.113.42 - - [14/May/2026:10:00:03 +0000] "POST /api HTTP/1.1" 500 256
```

---

## Example Output

```
=== LogSentry Audit Report ===

Total Requests: 22
Total 404 Errors: 3
Total 5xx Errors: 5

Top Client IPs

1. 192.168.1.10 (8 requests)
2. 10.0.0.5 (5 requests)
3. 203.0.113.42 (5 requests)
```

---

## Concepts Practiced

- File Handling
- Buffered File Reading
- Regular Expressions
- Maps (Hash Maps)
- Structs
- Slices
- Custom Sorting
- CLI Applications
- Log Processing
- Data Aggregation

---

## Future Improvements

- Export audit reports to a text file
- Support Apache and Nginx log formats
- Endpoint analytics
- Status code distribution chart
- Suspicious IP detection
- Command-line flags
- Execution time metrics
- JSON report generation

---

## Learning Outcomes

This project helped reinforce:

- Go file handling
- Streaming large files efficiently
- Pattern matching using Regular Expressions
- Data aggregation using maps
- Structs and slices
- Sorting custom data
- Building command-line utilities in Go

---

## Author

**Trishanth Sai**

GitHub: https://github.com/Trishanthsai
