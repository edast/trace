package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"runtime/trace"
	"sync"

	"github.com/edast/tracer/utils"
)

// pass filename as first argument
func main() {

	f, err := os.Create("multiple.out")
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

	var wg sync.WaitGroup

	results := make(chan string, len(data))
	wg.Add(len(data))

	for _, d := range data {
		go func(str string) {
			defer wg.Done()
			hash := sha512.Sum512([]byte(str))
			results <- fmt.Sprintf("%s : %x", str, hash)
		}(d)
	}

	wg.Wait()
	close(results)

	result := make([]string, 0, len(data))
	for d := range results {
		result = append(result, d)
	}

	utils.WriteToFile(result)
}
