package tests

import (
	"testing"
	"Govid/govid"
)

func TestGetGlobalData(t *testing.T) {
	output, err := govid.GetGlobalData()

	if err != nil {
		t.Fatalf("error getting global data %+v", err)
	}

	if output == nil {
		t.Fatalf("unable to get global data")
	}
}