
```go
package main

import (
	"fmt"
	"github.com/1makarov/go-pricempire"
	"log"
)

func main() {
	client := pricempire.NewClient("APIKEY")
	items, err := client.GetAllItemsBySites("csgoempire")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(items)
}
```