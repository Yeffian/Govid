package tests

import (
	"github.com/YeffyCodeGit/Govid/govid"
	"testing"
)

func TestGetCountryData(t *testing.T) {
	country, err := govid.GetCountryData("USA")
	if err != nil {
		t.Fatalf("unable to get country %+v", err)
	}

	if country.Country != "USA" {
		t.Errorf("invalid country")
	}
}

func TestGetAllCountriesData(t *testing.T) {
	data, err := govid.GetAllCountriesData()

	if err != nil {
		t.Fatalf("unable to get all countries %+v", err)
	}

	if data[0].Country != "World" && data[len(data) - 1].Country != "China" {
		t.Fatalf("invalid data")
	}
}

func TestGetGlobalDataByLimit(t *testing.T) {
	output, err := govid.GetCountryDataByLimit(2)

	if err != nil {
		t.Fatalf("error getting global data by limit %+v", err)
	}

	if output[0].Country != "World" && output[2].Country == "India" {
		t.Fatalf("invalid data")
	}
}

func TestGetGlobalDatatByEndLimit(t *testing.T) {
	output, err := govid.GetCountryDataByEndLimit(3, 1)

	if err != nil { t.Fatalf("error getting global data by end and start limit %+v", err) }

	if output[0].Country != "USA" && output[3].Country == "Brazil" {
		t.Fatalf("invalid data")
	}
}