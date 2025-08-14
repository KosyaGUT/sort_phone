package main

import (
	"fmt"
	"time"

	"github.com/kosyagut/sortphone/internal/reader"
	"github.com/kosyagut/sortphone/internal/sorter"
	"github.com/kosyagut/sortphone/internal/writer"
)

func main() {
	start := time.Now()
	read := reader.ReadFile("phones.txt")
	sortedRead := sorter.SortPhone(read)
	writer.WriterFile(sortedRead)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
