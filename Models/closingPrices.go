package Models

type ClosingOutputMetaData struct {
	Symbol		string	`json:"Symbol"`
	Days		int		`json:"Days"`
	CloseHigh	float64	`json:"Closing High"`
	CloseLow	float64	`json:"Closing Low"`
	CloseAvg	float64	`json:"Closing Average`
	CloseMed	float64	`json:"Closing Median"`
}

type ClosingOutputData struct {
	Metadata	ClosingOutputMetaData	`json:"MetaData"`
	DailyClose	map[string]float64		`json:"Closing Prices"`
}
