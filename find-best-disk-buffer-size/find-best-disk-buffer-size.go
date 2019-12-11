package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"time"
)

func main() {
	lowestSize := int64(math.MaxInt64)
	lowestTime := int64(math.MaxInt64)

	for i := int64(1); i <= 65536; i++ {
		file, err := os.Open("moby-dick-1st-chapter.txt")

		if err != nil {
			panic(err)
		}

		start := time.Now().UnixNano()
		buffer := make([]byte, i)

		for {
			_, err := file.Read(buffer)

			if err != nil {
				if err == io.EOF {
					break
				}

				panic(err)
			}
		}

		elapsed := time.Now().UnixNano() - start
		file.Close()
		fmt.Printf("[%d] %d ns\n", i, elapsed)

		if elapsed < lowestTime {
			lowestTime = elapsed
			lowestSize = i
		}

		file = nil
		buffer = nil
		runtime.GC()
	}

	fmt.Printf("Fastest read with buffer size of %d bytes @ %d ns\n", lowestSize, lowestTime)
}
