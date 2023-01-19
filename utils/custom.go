
package utils

import (
	"io"
	"net/http"
	"net/url"
)

// Customize the response with the way that you want
// But remember that respose must be legit in according within http struct.
type Custom struct {
	Header     http.Header    // Header contents like content-type, accept, etc...
	StatusCode int16          //output default status code - 200. (This app only returns 200 if OK, if don't will return the specified configurated status code from errors.)