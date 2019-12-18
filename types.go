package splunk

import (
	"encoding/xml"
)

type SearchSubmitResult struct {
	XMLName xml.Name `xml:"response"`
	Sid     string   `xml:"sid"`
}
