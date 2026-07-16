# 🛡️ LogSentry

A lightweight intrusion detection and log monitoring tool built with Go.

LogSentry analyzes server logs to detect suspicious activity, identify blacklisted IPs, perform GeoIP lookups, send email alerts, and generate JSON reports.

## ✨ Features

- Parse Apache/Nginx access logs
- Detect suspicious IP activity
- Automatic IP blacklisting
- GeoIP location lookup
- Email alerts
- JSON report generation
- Top IP & endpoint statistics
- Real-time log monitoring

## 🛠️ Tech Stack

- Go
- MaxMind GeoLite2
- SMTP
- JSON

## 🚀 Run

```bash
git clone https://github.com/yourusername/LogSentry.git
cd LogSentry
go mod tidy
go run .
```

## 📂 Project Structure

```
main.go
parser.go
geoip.go
email.go
monitor.go
report.go
utils.go
models.go
```

## 🔮 Future Improvements

- REST API
- Web Dashboard
- Docker Support
- Slack/Discord Alerts

---

Made using Go.
