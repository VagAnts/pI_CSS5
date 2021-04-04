
package helpers

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFiles(path string, rChan chan []string) chan []string {