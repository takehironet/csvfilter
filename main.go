package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		inputDelimiter  = flag.String("i", ",", "Input Delimiter")
		outputDelimiter = flag.String("o", ",", "Output Delimiter")
		columns         = flag.String("c", "1", "Output Columns separated by a space. Example: \"1 3 4 2\"")
	)

	flag.Parse()

	var fp *os.File
	fp = os.Stdin

	var cols []int
	for _, v := range strings.Split(string(*columns), " ") {
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		if i >= 0 {
			cols = append(cols, int(i))
		}
	}
	reader := csv.NewReader(fp)
	reader.Comma = []rune(*inputDelimiter)[0]
	reader.LazyQuotes = true

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		cs := pickColumn(row, cols)
		{
			for k, v := range cs {
				if k == 0 {
					fmt.Printf("\"%s\"", v)
				} else {
					fmt.Printf("%s\"%s\"", *outputDelimiter, v)
				}
			}
			fmt.Print("\n")
		}

	}
}

func pickColumn(arr []string, idx []int) []string {
	arrLength := len(arr)
	colLength := len(idx)
	if colLength < 1 {
		return nil
	}
	var res []string
	for _, i := range idx {
		if i > arrLength { // means "(i - 1) > (arrLength - 1)". "i" is 1-origin number.
			res = append(res, "")
		} else {
			res = append(res, arr[i-1])
		}
	}
	return res
}
