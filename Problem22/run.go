package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	file, err := os.Open("Problem22/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return
	}

	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.ReplaceAll(line, "\"", "")

	data := strings.Split(line, ",")

	sort.Strings(data)

	result := 0

	for index, element := range data {
		value := 0
		for _, char := range element {
			value += stringValueOf(string(char))
		}
		result += value * (index + 1)
	}

	fmt.Printf("Result: %[1]d\n", result)
}

func stringValueOf(letter string) int {
	var a = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(a, letter) + 1
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}