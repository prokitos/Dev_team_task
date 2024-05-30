package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"module/internal/models"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2/log"
)

// отправка гет запроса на указанный сайт с указанным параметром
func sendRequestToGet(first string, second string) float64 {

	var curUrl string = fmt.Sprintf("https://v6.exchangerate-api.com/v6/436f6e37d6cce64bd67bf475/pair/%s/%s", first, second)

	baseURL, _ := url.Parse(curUrl)
	params := url.Values{}
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Debug("Error connecting to external api")
		log.Error("Error getting values")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var results string = string(body)

	return ResultComputing(results)
}

func ResultComputing(jsonString string) float64 {

	var instanse models.MessageResult
	data := []byte(jsonString)
	err := json.Unmarshal(data, &instanse)
	if err != nil {
		log.Error("error in getting the values")
		log.Debug("error when convert currency into json file: " + jsonString)
	}

	return instanse.Conversion_rate
}
