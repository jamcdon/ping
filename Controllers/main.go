package Controllers

import (
	"net/http"
	"fmt"
	"io" // "io/ioutil"
	"encoding/json"
	
	"github.com/jamcdon/ping/Models"

	"github.com/gin-gonic/gin"
)

const apiKey = "FN8UEDB4KX3YR0EJ"

func readDaily(apiKey string, symbol string) Models.AVTimeSeriesDailyOutput {
	url = fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", apiKey, symbol)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var avOutput Models.AVTimeSeriesDailyOutput
	err = json.Unmarshal(body, avOutput)
	if err != nil(
		log.Fatal(err)
	)

	return avOutput
}

func Daily(c *gin.Context) {
	symbol = c.Param(symbol)
	TSDOutput, err := readDaily(apiKey, symbol)
	if err != nil (
		c.String(http.statusInternalServerError, "an error occured")
	)
	c.JSON(http.statusOK, TSDOutput)
}