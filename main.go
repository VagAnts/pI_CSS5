
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
