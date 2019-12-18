package splunk

import (
	"fmt"
	"net/url"
)

func (conn SplunkConnection) Search(search string) (string, error) {
	data := make(url.Values)
	data.Add("search", search)
	response, err := conn.httpPost(fmt.Sprintf("%s/services/search/jobs", conn.BaseURL), &data)
	return response, err
}
