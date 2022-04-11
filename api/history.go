package handler

import (
	"encoding/json"
	"fmt"
	"jds-weather/apiutils"
	"log"
	"net/http"
)

type TypeHistoaryResponse struct {
	Location apiutils.TypeLocation `json:"location"`
	Forecast struct {
		Forecastday []apiutils.TypeForecastday `json:"forecastday"`
	} `json:"forecast"`
}

func History(w http.ResponseWriter, r *http.Request) {
	url := apiutils.GetQuryUrl("history", r.URL.String())
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

	var history TypeHistoaryResponse
	err = json.NewDecoder(res.Body).Decode(&history)
	if err != nil {
		panic(err)
	}

	jsonResp, err := json.Marshal(history)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
