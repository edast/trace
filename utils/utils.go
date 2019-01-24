package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ReadFile reads file slice of strings
// will calculate sha512 hash for each line and will write result to /tmp/output.XXXXX.txt file
func ReadFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

// WriteToFile writes slice of strings to tmp file
func WriteToFile(data []string) {
	tmpfile, err := ioutil.TempFile("/tmp", "output.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()
	fmt.Println("output file:", tmpfile.Name())

	for _, d := range data {
		tmpfile.WriteString(d + "\n")
	}
}
