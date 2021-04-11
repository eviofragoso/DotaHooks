package gsi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaacov/observer/observer"
)

var obs *observer.Observer

// Listen for Dota client calls
func Listen(port int16, o *observer.Observer) error {
	obs = o
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	return nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	handleGameEvents(body)
}

func handleGameEvents(jsonString []byte) {
	gameState := parseGameState(jsonString)

	handleRoshanKilled(&gameState)
	handleTeamWiped(&gameState)
	handleSmokeGank(&gameState)
}

func handleRoshanKilled(gameState *GameState) {
	// fmt.Println("RoshanKilled")
}

func handleTeamWiped(gameState *GameState) {
	if gameState.Hero.Dire.Wiped {
		obs.Emit("dire:teamWiped")
	}

	if gameState.Hero.Radiant.Wiped {
		obs.Emit("radiant:teamWiped")
	}
}

func handleSmokeGank(gameState *GameState) {
	if gameState.Hero.Dire.SmokeGank {
		obs.Emit("dire:smokeGank")
	}

	if gameState.Hero.Radiant.SmokeGank {
		obs.Emit("radiant:smokeGank")
	}
}
