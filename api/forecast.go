package handler

import (
	"encoding/json"
	"fmt"
	"jds-weather/apiutils"
	"log"
	"net/http"
)

type TypeForecastResponse struct {
	Location apiutils.TypeLocation `json:"location"`
	Current  apiutils.TypeCurrent  `json:"current"`
	Forecast struct {
		Forecastday []apiutils.TypeForecastday `json:"forecastday"`
	} `json:"forecast"`
}

func Forecast(w http.ResponseWriter, r *http.Request) {
	// err := godotenv.Load("./.env")
	// if err != nil {
	// 	log.Fatalf("Cant read env file. Err: %s", err)
	// }

	url := apiutils.GetQuryUrl("forecast", r.URL.String())
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

	var forecast TypeForecastResponse

	err = json.NewDecoder(res.Body).Decode(&forecast)
	if err != nil {
		panic(err)
	}

	jsonResp, err := json.Marshal(forecast)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
