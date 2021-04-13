package main

import (
	"fmt"

	"github.com/eviofragoso/DotaHooks/internal/gsi"
	"github.com/yaacov/observer/observer"
)

func main() {
	events := gsi.Events{}
	o := observer.Observer{}
	o.Open()
	defer o.Close()

	o.AddListener(func(e interface{}) {
		switch e.(string) {
		case "dire:teamWiped":
			fmt.Println("Dire was Team wiped")
		case "radiant:teamWiped":
			fmt.Println("Radiant was Team wiped")
		case "dire:smokeGank":
			fmt.Println("Dire is prepared to fight")
		case "radiant:smokeGank":
			fmt.Println("Radiant is prepared to fight")
		}
	})

	gsi.Listen(3000, &o, &events)
}
