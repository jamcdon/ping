package Controllers

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"os"
	
	"github.com/jamcdon/ping/Models"

	"github.com/gin-gonic/gin"
)

var apiKey = os.Getenv("APIKEY")

func readDaily(apiKey string, symbol string) (Models.AVTimeSeriesDailyOutput, error) {
	var avOutput Models.AVTimeSeriesDailyOutput
	url := fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", apiKey, symbol)

	// production todo - error out when given api key error
	// alphavantage returns a 200 on api key error so unable to filter by response code
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return avOutput, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return avOutput, err
	}
	err = json.Unmarshal(body, &avOutput)
	if err != nil {
		log.Println(err)
		return avOutput, err
	}

	return avOutput, nil
}

func Daily(c *gin.Context) {
	symbol := c.Param("symbol")
	nDays := c.Param("days")
	nDaysInt, err := strconv.Atoi(nDays)
	if err != nil || nDaysInt < 0 {
		c.String(http.StatusInternalServerError, "An error occured. Please check your days parameter")
	}

	TSDOutput, err := readDaily(apiKey, symbol)
	if err != nil {
		c.String(http.StatusInternalServerError, "an error occured")
	}


	// map of closing price for set amount of days requested
	reducedTimeSeries := make(map[string]float64)
	// map of all data for reduced days - used for metadata calculation
	timeSeriesReplacement := make(map[string]Models.AVTSDDaily)

	// create slice of TimeSeries map keys for sorting
	var timeSeriesKeys []string
	for k, _ := range TSDOutput.TimeSeries {
		timeSeriesKeys = append(timeSeriesKeys, k)
	}
	sort.Strings(timeSeriesKeys)

	// set start position and check for failing conditions
	indexStart := len(timeSeriesKeys) - nDaysInt
	if indexStart < 0 {
		indexStart = 0
	}
	if indexStart > len(timeSeriesKeys){
		c.String(http.StatusInternalServerError, "an error occured")
	}
	timeSeriesKeys = timeSeriesKeys[indexStart:]

	// use timeSeriesKeys to rebuild shortened list of security/closing data
	for _, v := range timeSeriesKeys{
		reducedTimeSeries[v], err = strconv.ParseFloat(TSDOutput.TimeSeries[v].Close, 8)
		timeSeriesReplacement[v] = TSDOutput.TimeSeries[v]

		if err != nil {
			c.String(http.StatusInternalServerError, "an error occured")
		}
	}

	// replace TimeSeries data to reduced list for metadata calculations
	TSDOutput.TimeSeries = timeSeriesReplacement
	metadata, err := getMetaData(TSDOutput)
	if err != nil {
		c.String(http.StatusInternalServerError, "an error occured preparing metadata")
	}

	var outputObject Models.ClosingOutputData

	outputObject.Metadata = metadata
	outputObject.DailyClose = reducedTimeSeries

	c.JSON(http.StatusOK, outputObject)
}



func getMetaData(AVTSDOutput Models.AVTimeSeriesDailyOutput) (Models.ClosingOutputMetaData, error) {
	var closingValues []float64
	for _, v := range AVTSDOutput.TimeSeries {
		close, _ := strconv.ParseFloat(v.Close, 8)
		closingValues = append(closingValues, close)
	}
	if len(closingValues) <= 0 {
		return Models.ClosingOutputMetaData{}, nil
	}

	sort.Float64s(closingValues)

	high := closingValues[len(closingValues)-1]
	low := closingValues[0]
	median := findMedian(closingValues)
	average := findAverage(closingValues)

	metadata := Models.ClosingOutputMetaData{
		Symbol: AVTSDOutput.Metadata.Symbol,
		Days: len(closingValues),
		CloseHigh: high,
		CloseLow: low,
		CloseAvg: average,
		CloseMed: median,
	}
	
	return metadata, nil

}

func findMedian(values []float64) float64 {
	var median float64

	i := len(values) / 2
	if len(values) % 2 == 1 {
		median = values[i]
	} else {
		median = (values[i-1] + values[i]) / 2
	}
	return median
}

func findAverage(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	average := sum / float64(len(values))
	return average
}