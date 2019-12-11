package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"time"
)

func main() {
	from := int64(1)
	to := int64(65536)
	step := int64(1)

	flag.Int64Var(&from, "from", from, "First buffer size")
	flag.Int64Var(&to, "to", to, "Last buffer size")
	flag.Int64Var(&step, "step", step, "Step size")
	flag.Parse()

	lowestSize := int64(math.MaxInt64)
	lowestTime := int64(math.MaxInt64)
	file, err := os.Open("moby-dick-1st-chapter.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for i := from; i <= to; i += step {
		_, err := file.Seek(0, 0)

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
		fmt.Printf("[%d] %d ns\n", i, elapsed)

		if elapsed < lowestTime {
			lowestTime = elapsed
			lowestSize = i
		}

		buffer = nil
		runtime.GC()
	}

	fmt.Printf("Fastest read with buffer size of %d bytes @ %d ns\n", lowestSize, lowestTime)
}
