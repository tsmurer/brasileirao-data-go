package scraper

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(requestUrl string) (*http.Response, error) {
	res, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println("error: ")
		return nil, fmt.Errorf("makeWikiRequest: %v", err)
	}
	fmt.Println("response: ", res)
	return res, err
}

func GetBodyFromResponse(response *http.Response) (string, error) {
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("getResponseBody: %v", err)
	}
	return string(resBody), err
}
