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
	buffer  *bufio.Writer
	writer  io.Writer
)

func main() {
	defer flush()
	initWordScanner()
	initBufWriter()

	n := readInt()
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, readInt())
	}

	for i := 1; i < n; i++ {
		d := a[i] - a[i-1]
		switch {
		case d == 0:
			println("stay")
		case d > 0:
			printf("up %d\n", d)
		case d < 0:
			printf("down %d\n", abs(d))
		}
	}
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(writer, args...)
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(writer, f, args...)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readInt() int {
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func initLineScanner() {
	scanner = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1e4)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
}

func initWordScanner() {
	initLineScanner()
	scanner.Split(bufio.ScanWords)
}

func initBufWriter() {
	buffer = bufio.NewWriter(os.Stdout)
	writer = buffer
}

func flush() {
	if writer != nil {
		buffer.Flush()
	}
}
