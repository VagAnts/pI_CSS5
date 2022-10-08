
package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/leoff00/gocheckitout/api"
	"github.com/leoff00/gocheckitout/helpers"
	"github.com/leoff00/gocheckitout/logs"
)

func readOptionCode() int16 {
	var code int16
	fmt.Scan(&code)
	return code
}

var wg sync.WaitGroup

func main() {
	logs.WelcomeLog()

	res := make(chan []string)
	go helpers.ReadFiles("dummy.txt", res)
	urls := <-res

	readedCode := readOptionCode()

	switch readedCode {
	case 1:
		for _, url := range urls {
			wg.Add(1)