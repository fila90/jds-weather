package apiutils

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func GetQueryParamsFromUrl(input_url string) url.Values {
	u, err := url.Parse(input_url)
	if err != nil {
		log.Fatal("can't parse url params", err)
	}

	return u.Query()
}

func GetQuryUrl(queryUrl, clientUrl string) string {
	apiUrl := os.Getenv("VITE_API_URL")
	apiKey := os.Getenv("VITE_API_KEY")

	urlParams := GetQueryParamsFromUrl(clientUrl)

	newUrl := fmt.Sprintf("%s/%s.json?key=%s", apiUrl, queryUrl, apiKey)

	// loop params and append each one
	for param, paramVal := range urlParams {
		newUrl += "&" + param + "=" + paramVal[0]
	}

	return newUrl
}
