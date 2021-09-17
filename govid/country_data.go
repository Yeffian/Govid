package govid

// Stores covid data about any country
type CountryData struct {
	Country             string `json:"country"`
	Cases               int    `json:"cases"`
	TodayCases          int    `json:"todayCases"`
	Deaths              int    `json:"deaths"`
	TodayDeaths         int    `json:"todayDeaths"`
	Recovered           int    `json:"recovered"`
	Active              int    `json:"active"`
	Critical            int    `json:"critical"`
	CasesPerOneMillion  int    `json:"casesPerOneMillion"`
	DeathsPerOneMillion int    `json:"deathsPerOneMillion"`
	TotalTests          int    `json:"totalTests"`
	TestsPerOneMillion  int    `json:"testsPerOneMillion"`
}
