package main
import(
	"regexp"
)
var pattern =regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[.*?\]\s+"\S+\s+(\S+)\s+\S+"\s+(\d{3})`)
