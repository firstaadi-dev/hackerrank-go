package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'RemainderSorting' function below (and other code for sorting if needed).
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

func RemainderSorting(strArr []string) []string {
	var key []int
	result := make([]string, 0)
	strMap := make(map[int][]string)
	for _, v := range strArr {
		lenMod := len(v) % 3
		if _, ok := strMap[lenMod]; !ok {
			strMap[lenMod] = []string{v}
		} else {
			strMap[lenMod] = append(strMap[lenMod], v)
		}
	}

	for _, v := range strMap {
		if len(v) > 1 {
			sort.Strings(v)
		}
	}

	for k := range strMap {
		key = append(key, k)
	}
	sortedKey := sort.IntSlice(key)
	for k := range sortedKey {
		result = append(result, strMap[k]...)
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strArr []string

	for i := 0; i < int(strArrCount); i++ {
		strArrItem := readLine(reader)
		strArr = append(strArr, strArrItem)
	}

	result := RemainderSorting(strArr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
