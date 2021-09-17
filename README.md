# Govid

A simple Go library which lets you get information on Covid-19

## Examples

### Getting total data about all countires:

```go
package main

import (
	"Govid/govid"
	"fmt"
)

func main() {
	data, err := govid.GetGlobalData()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", data)
}
```

### Getting data about a specific country:

```go
package main

import (
	"Govid/govid"
	"fmt"
)

func main() {
	data, err := govid.GetCountryData("Bangladesh")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", data)
}
```

## Getting data about all the countries in the library

```go
package main

import (
	"Govid/govid"
	"fmt"
)

func main() {
	data, err := govid.GetAllCountriesData()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", data)
}
```