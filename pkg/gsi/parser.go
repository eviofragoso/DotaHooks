package gsi

type GameState struct {
	Previously *GameState `json:"previously"`
	Added      *GameState `json:"added"`
	Items      Items      `json:"items"`
	Map        Map        `json:"map"`
	Buildings  struct {
		Radiant interface{} `json:"radiant"`
		Dire    interface{} `json:"dire"`
	} `json:"building"`
	Draft  interface{} `json:"draft"`
	Hero   interface{} `json:"hero"`
	Player interface{} `json:"player"`
}

type Map struct {
	Name              string `json:"name"`
	MatchID           int16  `json:"matchid"`
	Gametime          string `json:"gametime"`
	Clocktime         string `json:"clocktime"`
	Daytime           bool   `json:"daytime"`
	NightstalkerNight bool   `json:"nightstalker_night"`
	Gamestate         string `json:"gamestate"`
	WinTeam           string `json:"win_team"`
	CustomGameName    string `json:"customgamename"`
}

type Items struct {
	Radiant struct {
		Player0 Player `json:"player0"`
		Player1 Player `json:"player1"`
		Player2 Player `json:"player2"`
		Player3 Player `json:"player3"`
		Player4 Player `json:"player4"`
	} `json:"team2"`
	Dire struct {
		Player5 Player `json:"player5"`
		Player6 Player `json:"player6"`
		Player7 Player `json:"player7"`
		Player8 Player `json:"player8"`
		Player9 Player `json:"player9"`
	} `json:"team3"`
}

type Buildings struct {
}

type Player struct {
	Slot1 struct {
		Item Item
	} `json:"slot1"`
	Slot2 struct {
		Item Item
	} `json:"slot12"`
	Slot3 struct {
		Item Item
	} `json:"slot3"`
	Slot4 struct {
		Item Item
	} `json:"slot4"`
	Slot5 struct {
		Item Item
	} `json:"slot5"`
	Slot6 struct {
		Item Item
	} `json:"slot6"`
}

type Item struct {
	Purchaser bool `json:"purchaser"`
	Passive   bool `json:"passive"`
}
