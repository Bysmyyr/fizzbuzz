package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	counter := 0
	for {
		var sb strings.Builder

		for i := 0; i < 1000; i++ {
			sb.WriteString(strconv.Itoa((counter + 1)))
			sb.WriteString("\n")
			sb.WriteString(strconv.Itoa((counter + 2)))
			sb.WriteString("\nFizz\n")
			sb.WriteString(strconv.Itoa((counter + 4)))
			sb.WriteString("\nBuzz\nFizz\n")
			sb.WriteString(strconv.Itoa((counter + 7)))
			sb.WriteString("\n")
			sb.WriteString(strconv.Itoa((counter + 8)))
			sb.WriteString("\nFizz\nBuzz\n")
			sb.WriteString(strconv.Itoa((counter + 11)))
			sb.WriteString("\nFizz\n")
			sb.WriteString(strconv.Itoa((counter + 13)))
			sb.WriteString("\n")
			sb.WriteString(strconv.Itoa((counter + 14)))
			sb.WriteString("\nFizzBuzz\n")
		}

		os.Stdout.Write([]byte(sb.String()))
		counter += 15
	}

}

// fmt.Sprintf
// stringWriter
