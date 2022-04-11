package handler

import (
	"encoding/json"
	"fmt"
	"jds-weather/apiutils"
	"log"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	url := apiutils.GetQuryUrl("search", r.URL.String())
	fmt.Println(url)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("can't create new request", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	var searchlocations []apiutils.TypeSearch
	err = json.NewDecoder(res.Body).Decode(&searchlocations)
	if err != nil {
		panic(err)
	}

	jsonResp, err := json.Marshal(searchlocations)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
