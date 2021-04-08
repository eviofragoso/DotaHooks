package gsi

type GameState struct {
	Previously interface{} `json:"previously"`
	Added      struct {
		Items struct {
			Radiant struct {
				Player0 AddedPlayer `json:"player0"`
				Player1 AddedPlayer `json:"player1"`
				Player2 AddedPlayer `json:"player2"`
				Player3 AddedPlayer `json:"player3"`
				Player4 AddedPlayer `json:"player4"`
			} `json:"team2"`
			Dire struct {
				Player5 AddedPlayer `json:"player5"`
				Player6 AddedPlayer `json:"player6"`
				Player7 AddedPlayer `json:"player7"`
				Player8 AddedPlayer `json:"player8"`
				Player9 AddedPlayer `json:"player9"`
			} `json:"team3"`
		} `json:"items"`
	} `json:"added"`
	Items     struct{} `json:"items"`
	Buildings struct {
		Radiant interface{} `json:"radiant"`
		Dire    interface{} `json:"dire"`
	} `json:"building"`
	Draft  interface{} `json:"draft"`
	Hero   interface{} `json:"hero"`
	Player interface{} `json:"player"`
}

type AddedPlayer struct {
	Slot1 struct {
		AddedItem AddedItem
	} `json:"slot1"`
	Slot2 struct {
		AddedItem AddedItem
	} `json:"slot12"`
	Slot3 struct {
		AddedItem AddedItem
	} `json:"slot3"`
	Slot4 struct {
		AddedItem AddedItem
	} `json:"slot4"`
	Slot5 struct {
		AddedItem AddedItem
	} `json:"slot5"`
	Slot6 struct {
		AddedItem AddedItem
	} `json:"slot6"`
}

type AddedItem struct {
	Purchaser bool `json:"purchaser"`
	Passive   bool `json:"passive"`
}
