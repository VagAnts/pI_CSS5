
package api

import (
	"fmt"

	"github.com/leoff00/gocheckitout/helpers"
)

//Make request returning only the url and the status code.
func MakeRequest(url string) error {