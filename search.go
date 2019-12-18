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
	fmt.Println("attempting unmarshalling")
	err = xml.Unmarshal([]byte(response), &searchResponse)
	if err != nil {
		fmt.Println("got an error")
	} else {
		fmt.Println("no error")
	}
	return searchResponse.Sid, err
}
