package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var (
    scanner *bufio.Scanner
    buffer  *bufio.Writer
    writer  io.Writer
)

func main() {
    defer flush()
    initWordScanner()
    initBufWriter()
    A := make([]int, 6)
    for i := 0; i < 6; i++ {
        A[i] = readInt()
    }
    sort.Ints(A)
    printf("%d", A[3])
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(writer, f, args...)
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
	buf := make([]byte, 6)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
}

func flush() {
    if writer != nil {
        buffer.Flush()
    }
}

