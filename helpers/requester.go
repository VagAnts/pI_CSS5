package helpers

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	utils "github.com/leoff00/gocheckitout/utils"
)

//Helper to make request with options and custom return.
func Requester(url string) (*utils.Custom, error) 