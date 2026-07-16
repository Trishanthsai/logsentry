package main
import(
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)
func monitorlogs(filename string){
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.Seek(0, io.SeekEnd) 
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)	
	}
	for {
    	for scanner.Scan() {
       	 	line := scanner.Text()
        	fmt.Println(line)
    	}
		time.Sleep(2 * time.Second)
	}
}