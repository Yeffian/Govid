package govid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// The different endpoints of the API
const (
	AllDataEndpoint = "https://coronavirus-19-api.herokuapp.com/all"
	AllCountriesEndpoint = "https://coronavirus-19-api.herokuapp.com/countries"
	CountryDataEndpoint = "https://coronavirus-19-api.herokuapp.com/countries/"
)

// Gets data about the global cases, recoveries and deaths
// Returns:
//   (*GlobalData, error)
func GetGlobalData() (*GlobalData, error) {
	res, err := http.Get(AllDataEndpoint)

	// Clean up the request
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil { return nil, err }

	body, err := ioutil.ReadAll(res.Body)

	if err != nil { return nil, err }

	var data GlobalData
	err = json.Unmarshal(body, &data)

	if err != nil { return nil, err }

	return &data, nil
}

// Gets data about a given country.
// Params:
//  country string
// Returns:
//  (*CountryData, error)
func GetCountryData(country string) (*CountryData, error) {
	res, err := http.Get(CountryDataEndpoint + country)

	if err != nil { return nil, err }

	// Clean up the request
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil { return nil, err }

	var data CountryData
	err = json.Unmarshal(body, &data)

	if err != nil { return nil, err }

	return &data, err
}

// Gets a list which contains all the data about every available country
// Returns:
//  ([]CountryData, error)
func GetAllCountriesData() ([]CountryData, error) {
	res, err := http.Get(AllCountriesEndpoint)

	if err != nil { return nil, err }

	// Clean up the request
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil { return nil, err }

	var data []CountryData

	err = json.Unmarshal(body, &data)

	if err != nil { return nil, err }

	return data, err
}

// The same as GetAllCountriesData, but allows you to explicitly specify a limit on many elements
// to get.
// Params:
//    limit int
// Returns:
//    ([]CountryData, error)
func GetCountriesDataByLimit(limit int) ([]CountryData, error) {
	arr, err := GetAllCountriesData()

	if err != nil { return nil, err }

	return arr[0:limit], nil
}


// Similar to GetCountriesDataByLimit, but lets you specify a start and end limit for more control.
// Params:
//    endLimit int
//    startLimit int
// Returns:
//   ([]CountryData, error)
func GetCountriesDataByEndLimit(endLimit int, startLimit int) ([]CountryData, error) {
	arr, err := GetAllCountriesData()

	if err != nil { return nil, err }

	return arr[startLimit:endLimit], nil
}