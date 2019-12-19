package splunk

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

func (conn SplunkConnection) Search(search string) (string, error) {
	data := make(url.Values)
	data.Add("search", search)
	response, err := conn.httpPost(fmt.Sprintf("%s/services/search/jobs", conn.BaseURL), &data)
	var searchResponse SearchSubmitResult
	xml.Unmarshal([]byte(response), &searchResponse)
	return searchResponse.Sid, err
}
