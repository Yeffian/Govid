package main

import (
	"Govid/govid"
	"fmt"
)

func main() {
	out, err := govid.GetGlobalData()

	if err != nil { panic(err) }

	fmt.Printf("%+v", out)
	fmt.Println("")

	output, err := govid.GetCountryData("USA")

	if err != nil { panic(err) }

	fmt.Printf("%+v", output)
	fmt.Println("")

	// Lucky there isnt more or these names would have gotten wild
	outputvalue, err := govid.GetAllCountriesData()

	if err != nil { panic(err) }

	fmt.Printf("%+v", outputvalue)
}
