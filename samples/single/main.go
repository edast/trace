package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"runtime/trace"

	"github.com/edast/tracer/utils"
)

// pass filename as first argument
// will calculate sha512 hash for each line and will write result to /tmp/output.XXXXX.txt file
func main() {

	f, err := os.Create("single.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	data := utils.ReadFile(os.Args[1])

	result := make([]string, 0, len(data))
	for _, d := range data {
		hash := sha512.Sum512([]byte(d))
		result = append(result, fmt.Sprintf("%s : %x", d, hash))
	}

	utils.WriteToFile(result)
}
