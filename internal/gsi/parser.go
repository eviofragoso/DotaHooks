package gsi

import (
	"encoding/json"
	"time"
)

// Bool custom boolean type

// UnmarshalJSON override of json marshal to handle empty booleans
// func (value *heroPlayerData) UnmarshalJSON(b []byte) (err error) {
// 	var hpd map[string]interface{}
// 	json.Unmarshal(b, &hpd)

// 	value.Alive = bool(hpd["alive"] == true)
// 	value.Dead = bool(hpd["alive"] == false)
// 	value.Smoked = bool(hpd["smoked"] == true)

// 	return err
// }

// GameState Data structure that holds data received from the client
type GameState struct {
	Previously *GameState   `json:"previously"`
	Added      *GameState   `json:"added"`
	Items      *items       `json:"items"`
	Map        *mapData     `json:"map"`
	Buildings  *buildings   `json:"buildings"`
	Player     *playerStats `json:"player"`
	Hero       *hero        `json:"hero"`
	Events     *Events
}

// Events holder for keep state of notifications
type Events struct {
	aegisTaken           time.Time
	teamWipedTimeDire    time.Time
	teamWipedTimeRadiant time.Time
	smokeGankDire        time.Time
	smokeGankRadiant     time.Time
}

type mapData struct {
	Name               string `json:"name,omitempty"`
	MatchID            int16  `json:"matchid,omitempty"`
	Gametime           string `json:"game_time,omitempty"`
	Clocktime          int8   `json:"clock_time,omitempty"`
	Daytime            bool   `json:"daytime,omitempty"`
	NightstalkerNight  bool   `json:"nightstalker_night,omitempty"`
	Paused             bool   `json:"paused,omitempty"`
	WinTeam            string `json:"win_team,omitempty"`
	CustomGameName     string `json:"customgamename,omitempty"`
	RoshanState        string `json:"roshan_state,omitempty"`
	RoshanRevivalTimer int8   `json:"roshan_state_end_seconds,omitempty"`
}

type barracks struct {
	Health    int16 `json:"health,omitempty"`
	MaxHealth int16 `json:"max_health,omitempty"`
}

type buildings struct {
	Radiant radiantBuildings `json:"radiant"`
	Dire    direBuildings    `json:"dire"`
}

type radiantBuildings struct {
	RangedMiddleBarracks *barracks `json:"good_rax_range_mid"`
	RangedBotBarracks    *barracks `json:"good_rax_range_bot"`
	RangedTopBarracks    *barracks `json:"good_rax_range_top"`
	MeleeMiddleBarracks  *barracks `json:"good_rax_melee_mid"`
	MeleedBotBarracks    *barracks `json:"good_rax_melee_bot"`
	MeleedTopBarracks    *barracks `json:"good_rax_melee_top"`
}

type direBuildings struct {
	RangedMiddleBarracks *barracks `json:"bad_rax_range_mid"`
	RangedBotBarracks    *barracks `json:"bad_rax_range_bot"`
	RangedTopBarracks    *barracks `json:"bad_rax_range_top"`
	MeleeMiddleBarracks  *barracks `json:"bad_rax_melee_mid"`
	MeleedBotBarracks    *barracks `json:"bad_rax_melee_bot"`
}

type player struct {
	SteamID        string `json:"steamid,omitempty"`
	Name           string `json:"name,omitempty"`
	Kills          int8   `json:"kills,omitempty"`
	Deaths         int8   `json:"deaths,omitempty"`
	Assists        int8   `json:"assists,omitempty"`
	LastHits       int8   `json:"last_hits,omitempty"`
	KillStreak     int8   `json:"kill_streak,omitempty"`
	TeamName       string `json:"team_name,omitempty"`
	GPM            int8   `json:"gpm,omitempty"`
	XPM            int8   `json:"xpm,omitempty"`
	NetWorth       int16  `json:"net_worth,omitempty"`
	HeroDamage     int16  `json:"hero_damage,omitempty"`
	WardsPurchased int8   `json:"wards_purchased,omitempty"`
	WardsPlaced    int8   `json:"wards_placed,omitempty"`
	WardsDestroyed int8   `json:"wards_destroyed,omitempty"`
	RunesActivated int8   `json:"runes_activated,omitempty"`
	CampsStacked   int8   `json:"camps_stacked,omitempty"`
}

type playerStats struct {
	Radiant *playerStatsRadiant `json:"team2"`
	Dire    *playerStatsDire    `json:"team3"`
}

type playerStatsRadiant struct {
	Player0 *player `json:"player0"`
	Player1 *player `json:"player1"`
	Player2 *player `json:"player2"`
	Player3 *player `json:"player3"`
	Player4 *player `json:"player4"`
}

type playerStatsDire struct {
	Player5 *player `json:"player5"`
	Player6 *player `json:"player6"`
	Player7 *player `json:"player7"`
	Player8 *player `json:"player8"`
	Player9 *player `json:"player9"`
}

type item struct {
	Purchaser int8   `json:"purchaser,omitempty"`
	Passive   bool   `json:"passive,omitempty"`
	Name      string `json:"name,omitempty"`
}

type playerItems struct {
	Slot1  *item `json:"slot1,omitempty"`
	Slot2  *item `json:"slot2,omitempty"`
	Slot3  *item `json:"slot3,omitempty"`
	Slot4  *item `json:"slot4,omitempty"`
	Slot5  *item `json:"slot5,omitempty"`
	Slot6  *item `json:"slot6,omitempty"`
	Slot7  *item `json:"slot7,omitempty"`
	Slot8  *item `json:"slot8,omitempty"`
	Stash0 *item `json:"stash0,omitempty"`
	Stash1 *item `json:"stash1,omitempty"`
	Stash2 *item `json:"stash2,omitempty"`
	Stash3 *item `json:"stash3,omitempty"`
	Stash4 *item `json:"stash4,omitempty"`
	Stash5 *item `json:"stash5,omitempty"`
}

type items struct {
	Radiant *itemsRadiant `json:"team2"`
	Dire    *itemsDire    `json:"team3"`
}

type itemsRadiant struct {
	AegisTaken bool
	Player0    *playerItems `json:"player0"`
	Player1    *playerItems `json:"player1"`
	Player2    *playerItems `json:"player2"`
	Player3    *playerItems `json:"player3"`
	Player4    *playerItems `json:"player4"`
}

type itemsDire struct {
	AegisTaken bool
	Player5    *playerItems `json:"player5"`
	Player6    *playerItems `json:"player6"`
	Player7    *playerItems `json:"player7"`
	Player8    *playerItems `json:"player8"`
	Player9    *playerItems `json:"player9"`
}

type heroPlayerData struct {
	Alive  bool `json:"alive,omitempty"`
	Dead   bool
	Smoked bool `json:"smoked,omitempty"`
}

type heroTeamDataRadiant struct {
	SmokeGank bool
	Wiped     bool
	Player0   *heroPlayerData `json:"player0"`
	Player1   *heroPlayerData `json:"player1"`
	Player2   *heroPlayerData `json:"player2"`
	Player3   *heroPlayerData `json:"player3"`
	Player4   *heroPlayerData `json:"player4"`
}

type heroTeamDataDire struct {
	SmokeGank bool
	Wiped     bool
	Player5   *heroPlayerData `json:"player5"`
	Player6   *heroPlayerData `json:"player6"`
	Player7   *heroPlayerData `json:"player7"`
	Player8   *heroPlayerData `json:"player8"`
	Player9   *heroPlayerData `json:"player9"`
}

type hero struct {
	Radiant *heroTeamDataRadiant `json:"team2"`
	Dire    *heroTeamDataDire    `json:"team3"`
}

func parseGameState(jsonData []byte, events *Events) GameState {
	added := GameState{}
	previously := GameState{}

	gameState := GameState{
		Added:      &added,
		Previously: &previously,
		Events:     events,
	}

	json.Unmarshal(jsonData, &gameState)
	updateHeroData(&gameState)
	updateAegisTaken(&gameState)

	return gameState
}

func updateHeroData(gameState *GameState) {
	if gameState.Hero == nil {
		return
	}

	checkRadiantTeamWiped(gameState.Hero.Radiant)
	checkDireTeamWiped(gameState.Hero.Dire)

	checkSmokeGankRadiant(gameState.Hero.Radiant)
	checkSmokeGankDire(gameState.Hero.Dire)
}

func updateAegisTaken(gameState *GameState) {
	if gameState.Map.RoshanState == "alive" &&
		(gameState.Items.Dire.AegisTaken || gameState.Items.Radiant.AegisTaken) {

		gameState.Items.Dire.AegisTaken = false
		gameState.Items.Radiant.AegisTaken = false

		return
	}

	if gameState.Map.RoshanState == "respawn_base" {
		updateAegisDire(gameState.Items.Dire)
		checkAegisRadiant(gameState.Items.Radiant)
	}
}

func updateAegisDire(items *itemsDire) {
	players := [5]*playerItems{
		items.Player5,
		items.Player6,
		items.Player7,
		items.Player8,
		items.Player9,
	}

	for i := 0; i < len(players); i++ {
		playerItems := [6]*item{
			players[i].Slot1,
			players[i].Slot2,
			players[i].Slot3,
			players[i].Slot4,
			players[i].Slot5,
			players[i].Slot6,
		}

		for j := 0; j < len(playerItems); j++ {
			if playerItems[j] != nil && playerItems[j].Name == "item_aegis" {
				items.AegisTaken = true
				return
			}
		}
	}

	items.AegisTaken = false
}

func checkAegisRadiant(items *itemsRadiant) {
	players := [5]*playerItems{
		items.Player0,
		items.Player1,
		items.Player2,
		items.Player3,
		items.Player4,
	}

	for i := 0; i < len(players); i++ {
		playerItems := [6]*item{
			players[i].Slot1,
			players[i].Slot2,
			players[i].Slot3,
			players[i].Slot4,
			players[i].Slot5,
			players[i].Slot6,
		}

		for j := 0; j < len(playerItems); j++ {
			if playerItems[j] != nil && playerItems[j].Name == "item_aegis" {
				items.AegisTaken = true
				return
			}
		}
	}

	items.AegisTaken = false
}

func checkRadiantTeamWiped(data *heroTeamDataRadiant) {
	values := [5]bool{
		data.Player0.Alive,
		data.Player1.Alive,
		data.Player2.Alive,
		data.Player3.Alive,
		data.Player4.Alive,
	}

	for i := 0; i < len(values); i++ {
		if values[i] == true {
			data.Wiped = false
			return
		}
	}

	data.Wiped = true
}

func checkDireTeamWiped(data *heroTeamDataDire) {
	values := [5]bool{
		data.Player5.Alive,
		data.Player6.Alive,
		data.Player7.Alive,
		data.Player8.Alive,
		data.Player9.Alive,
	}

	for i := 0; i < len(values); i++ {
		if values[i] == true {
			data.Wiped = false
			return
		}
	}

	data.Wiped = true
}

func checkSmokeGankRadiant(data *heroTeamDataRadiant) {
	values := [5]bool{
		data.Player0.Smoked,
		data.Player1.Smoked,
		data.Player2.Smoked,
		data.Player3.Smoked,
		data.Player4.Smoked,
	}

	counter := 0

	for i := 0; i < len(values); i++ {
		if counter >= 2 {
			data.SmokeGank = true
			return
		}
		if values[i] {
			counter++
		}
	}
}

func checkSmokeGankDire(data *heroTeamDataDire) {
	values := [5]bool{
		data.Player5.Smoked,
		data.Player6.Smoked,
		data.Player7.Smoked,
		data.Player8.Smoked,
		data.Player9.Smoked,
	}

	counter := 0

	for i := 0; i < len(values); i++ {
		if counter >= 2 {
			data.SmokeGank = true
			return
		}
		if values[i] {
			counter++
		}
	}
}
