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

## Building from source

First, clone the repo:
```bash
git clone https://github.com/YeffyCodeGit/Govid
```

Run the unit tests:
```bash
go test ./govid/tests
```

## Credits
Thanks to @WaifuShork for helping me on this project
This project uses the Covid-19 API by Javier Aviles. You can find it here: https://github.com/javieraviles/covidAPI