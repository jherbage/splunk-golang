package splunk

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
)

func (conn SplunkConnection) Search(search string, timeout int) (string, error) {
	data := make(url.Values)
	data.Add("search", search)
	response, err := conn.httpPost(fmt.Sprintf("%s/services/search/jobs", conn.BaseURL), &data)
	var searchResponse SearchSubmitResult
	xml.Unmarshal([]byte(response), &searchResponse)
	// We will wait until timeout before giving up on the search
	for expired := time.Now().Add(time.Duration(timeout) * time.Second); expired.Sub(time.Now()) > 0; time.Sleep(5 * time.Second) {
		response, err = conn.httpGet(fmt.Sprintf("%s/services/search/jobs/"+searchResponse.Sid, conn.BaseURL), nil)
	}
	return response, err
}
