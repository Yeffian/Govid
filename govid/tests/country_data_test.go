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