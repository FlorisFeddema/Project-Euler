package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	data := [100][100]int{}
	lineNumber := 0

	file, err := os.Open("Problem67/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")

		if len(line) == 0 {
			break
		}

		list := strings.Split(line, " ")

		for i := 0; i < len(list); i++ {
			data[lineNumber][i], _ = strconv.Atoi(list[i])
		}
		lineNumber++
	}

	for i := len(data) - 2 ; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if data[i + 1][j] > data[i + 1][j + 1] {
				data[i][j] +=  data[i + 1][j]
			} else {
				data[i][j] += data[i + 1][j + 1]
			}
		}
	}

	result := data[0][0]
	fmt.Printf("Result: %[1]d\n", result)
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}