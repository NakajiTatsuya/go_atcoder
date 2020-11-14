package main
 
import (
	"bufio"
	"fmt"
	"os"
	"io"
	"sort"
)

var (
    scanner *bufio.Scanner
    buffer  *bufio.Writer
    writer  io.Writer
)

const (
	initialBufSize = 10e4
	maxBufSize     = 10e6
)

func main() {
	defer flush()
	initWordScanner()
	initBufWriter()
	line := readWord()
	words := make([]string, 0, 100000)
    splitLineIntoWords(line, &words)
    sort.Strings(words)

    for _, w := range words {
    	original := string(w[0] + 'A' - 'a') + w[1:len(w)-1] + string(w[len(w)-1] + 'A' - 'a')
        printf("%s", original)
    }
    printf("\n")
}

func splitLineIntoWords(w string, pwords *[]string) {
	startIndex := 0
    upperCaseCount := 0

	for i := 0; i < len(w); i++ {
		if 'A' <= w[i] && 'Z' >= w[i] {
			upperCaseCount += 1
		}
		if upperCaseCount == 2 {
			// インデックスがstartIndexからiまでが単語
			word := string(w[startIndex] - 'A' + 'a') + w[startIndex+1:i] + string(w[i] - 'A' + 'a')
			*pwords = append(*pwords, word)
			upperCaseCount = 0
			startIndex = i+1
		}
    }
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(writer, f, args...)
}

func readWord() string {
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func initBufWriter() {
    buffer = bufio.NewWriter(os.Stdout)
    writer = buffer
}

func initLineScanner() {
    scanner = bufio.NewScanner(os.Stdin)
    buf := make([]byte, initialBufSize)
    scanner.Buffer(buf, maxBufSize)
}

func initWordScanner() {
	initLineScanner()
	scanner.Split(bufio.ScanWords)
}

func flush() {
    if writer != nil {
        buffer.Flush()
    }
}
