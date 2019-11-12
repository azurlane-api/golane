[![Go Report Card](https://goreportcard.com/badge/github.com/KurozeroPB/golane)](https://goreportcard.com/report/github.com/KurozeroPB/golane)

# golane
Wrapper for the unofficial azur lane json api in Go

## Docs
https://godoc.org/github.com/KurozeroPB/golane

## Example
```go
package main

import (
	"fmt"
	"github.com/KurozeroPB/golane"
)

func main() {
	var azurlane = new(golane.AzurLane)
	azurlane.Init("custom_ua/v0.1.0")

	ships, err := azurlane.GetShips(golane.Order.RARITY, "Super Rare")
	if err != nil {
		fmt.Printf("Something bad happened:\n%s", err.Error())
		return
	}

	for i := 0; i < len(ships); i++ {
		fmt.Printf("[%s]: %s\n", ships[i].ID, ships[i].Name)
	}
}
```

## Support
![discord](https://discordapp.com/api/v6/guilds/240059867744698368/widget.png?style=banner2)