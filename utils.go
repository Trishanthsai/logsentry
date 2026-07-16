package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}
func addToBlacklist(ip string) {

	file, err := os.OpenFile("blacklist.txt",
		os.O_APPEND|os.O_WRONLY,
		0644)

	if err != nil {
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, ip)

	writer.Flush()
}