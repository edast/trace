package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"

	"github.com/edast/tracer/utils"
)

// pass filename as first argument
// will calculate sha512 hash for each line and will write result to /tmp/output.XXXXX.txt file
func main() {

	f, err := os.Create("pool.out")
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

	workerCount := runtime.NumCPU()
	fmt.Println("number of workers: ", workerCount)

	jobs := make(chan string, workerCount)
	results := make(chan string, len(data))

	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for str := range jobs {
				hash := sha512.Sum512([]byte(str))
				results <- fmt.Sprintf("%s : %x", str, hash)
			}
		}()
	}

	for _, d := range data {
		jobs <- d
	}
	close(jobs)
	wg.Wait()

	close(results)

	result := make([]string, 0, len(data))
	for d := range results {
		result = append(result, d)
	}

	utils.WriteToFile(result)
}
