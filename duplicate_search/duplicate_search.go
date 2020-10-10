package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
)

var (
    scanner *bufio.Scanner
    buffer *bufio.Writer
    writer io.Writer
)

func main() {
    defer flush()
    initWordScanner()
    initBufWriter()
    n := readInt()
    N := make([]int, n)
    exists := make([]int, n)

    for i := 0; i < n; i++ {
        N[i] = readInt()
    }
    for i := 0; i < n; i++ {
        exists[N[i]-1]++
    }
    var duplicate int = -1
    var original int = -1

    for i := 0; i < n; i++ {
      if exists[i] == 0 {
          original = i + 1
      } else if exists[i] == 2 {
          duplicate = i + 1
      }
    }

    if duplicate == -1 {
        printf("Correct")
    } else {
        printf("%d %d", duplicate, original)
    }
}

func readInt() int {
    scanner.Scan()
    n, err := strconv.Atoi(scanner.Text())
    if err != nil {
        panic(err)
    }
    return n
}

func initBufWriter() {
    buffer = bufio.NewWriter(os.Stdout)
    writer = buffer
}

func initWordScanner() {
    initLineScanner()
    scanner.Split(bufio.ScanWords)
}

func initLineScanner() {
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1e4), bufio.MaxScanTokenSize)
}



func flush() {
    if writer != nil {
        buffer.Flush()
    }
}

func printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(writer, format, args...)
}

