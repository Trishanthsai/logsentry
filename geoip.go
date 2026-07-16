package main
import(
	"net"
	"github.com/oschwald/geoip2-golang"
)


func getCity(ip string) string {
	db, err := geoip2.Open("data/GeoLite2-City.mmdb")
	if err != nil {
		return "Unknown"
	}
	defer db.Close()
	parsedIP := net.ParseIP(ip)
	record, err := db.City(parsedIP)
	if err != nil {
		return "Unknown"
	}
	if record.City.Names["en"] == "" {
		return record.Country.Names["en"]
	}
	return record.Country.Names["en"] + ", " + record.City.Names["en"]
}
