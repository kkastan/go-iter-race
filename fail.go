package main

import (
	"fmt"
	"sync"
)

func doPrint(val string) {
	fmt.Printf("%s\n", val)
}

func main() {
	var wg sync.WaitGroup
	tokens := []string{"Charlie", "Mac", "Dennis", "Dee", "Frank"}

	for _, token := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doPrint(token)
		}()
	}

	wg.Wait()
}
