package gsi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaacov/observer/observer"
)

var obs *observer.Observer
var evt *Events

// Listen for Dota client calls
func Listen(port int16, o *observer.Observer, events *Events) error {
	obs = o
	evt = events

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
	gameState := parseGameState(jsonString, evt)
	// s, _ := json.MarshalIndent(gameState.Added, "", "\t")
	// fmt.Print(string(s))

	handleAegisTaken(&gameState)
	handleTeamWiped(&gameState)
	handleSmokeGank(&gameState)
}

func handleAegisTaken(gameState *GameState) {
	// fmt.Println("RoshanKilled")
	var now time.Time = time.Now()

	if !gameState.Events.aegisTaken.IsZero() {
		if gameState.Map.RoshanState == "alive" {
			gameState.Events.aegisTaken = time.Time{}
		}

		return
	}

	if gameState.Items.Dire.AegisTaken {
		gameState.Events.aegisTaken = now
		obs.Emit("dire:aegisTaken")
	}

	if gameState.Items.Radiant.AegisTaken {
		gameState.Events.aegisTaken = now
		obs.Emit("radiant:aegisTaken")
	}
}

func handleTeamWiped(gameState *GameState) {
	var now time.Time = time.Now()
	if gameState.Events.teamWipedTimeDire.IsZero() ||
		gameState.Events.teamWipedTimeDire.Before(now.Add(-1*time.Minute)) {

		if gameState.Hero.Dire.Wiped {
			gameState.Events.teamWipedTimeDire = now
			obs.Emit("dire:teamWiped")
		}
	}

	if gameState.Events.teamWipedTimeRadiant.IsZero() ||
		gameState.Events.teamWipedTimeRadiant.Before(now.Add(-1*time.Minute)) {

		if gameState.Hero.Radiant.Wiped {
			gameState.Events.teamWipedTimeRadiant = now
			obs.Emit("radiant:teamWiped")
		}
	}
}

func handleSmokeGank(gameState *GameState) {
	var now time.Time = time.Now()

	if gameState.Events.smokeGankDire.IsZero() ||
		gameState.Events.smokeGankDire.Before(now.Add(-1*time.Minute)) {

		if gameState.Hero.Dire.SmokeGank {
			gameState.Events.smokeGankDire = now
			obs.Emit("dire:smokeGank")
		}
	}

	if gameState.Events.smokeGankRadiant.IsZero() ||
		gameState.Events.smokeGankRadiant.Before(now.Add(-1*time.Minute)) {

		if gameState.Hero.Radiant.SmokeGank {
			gameState.Events.smokeGankRadiant = now
			obs.Emit("radiant:smokeGank")
		}
	}
}
