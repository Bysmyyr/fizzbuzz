package main

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

const routines = 2

func main() {
	eightZeroLock := &sync.Mutex{}
	startLock := eightZeroLock
	endLock := &sync.Mutex{}
	endLock.Lock()
	for i := 0; i < routines-1; i++ {
		go routine(i, startLock, endLock)
		startLock = endLock
		endLock = &sync.Mutex{}
		endLock.Lock()
	}
	go routine(7, startLock, eightZeroLock)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func routine(num int, start, end *sync.Mutex) {
	counter := num * 15
	for {
		var sb strings.Builder

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

		start.Lock()
		os.Stdout.Write([]byte(sb.String()))
		end.Unlock()
		counter += 15 * routines
	}
}

// fmt.Sprintf
// stringWriter
