package gsi

type GameState struct {
	Previously *GameState  `json:"previously"`
	Added      *GameState  `json:"added"`
	Items      Items       `json:"items"`
	Map        Map         `json:"map"`
	Buildings  Buildings   `json:"buildings"`
	Player     PlayerStats `json:"player"`
}

type Map struct {
	Name               string `json:"name"`
	MatchID            int16  `json:"matchid"`
	Gametime           string `json:"game_time"`
	Clocktime          string `json:"clock_time"`
	Daytime            bool   `json:"daytime"`
	NightstalkerNight  bool   `json:"nightstalker_night"`
	Paused             bool   `json:"paused"`
	WinTeam            string `json:"win_team"`
	CustomGameName     string `json:"customgamename"`
	RoshanState        string `json:"roshan_state"`
	RoshanRevivalTimer int8   `json:"roshan_state_end_seconds"`
}

type Barracks struct {
	Health    int16 `json:"health"`
	MaxHealth int16 `json:"max_health"`
}

type Buildings struct {
	Radiant struct {
		RangedMiddleBarracks Barracks `json:"good_rax_range_mid"`
		RangedBotBarracks    Barracks `json:"good_rax_range_bot"`
		RangedTopBarracks    Barracks `json:"good_rax_range_top"`
		MeleeMiddleBarracks  Barracks `json:"good_rax_melee_mid"`
		MeleedBotBarracks    Barracks `json:"good_rax_melee_bot"`
		MeleedTopBarracks    Barracks `json:"good_rax_melee_top"`
	} `json:"radiant"`
	Dire struct {
		RangedMiddleBarracks Barracks `json:"bad_rax_range_mid"`
		RangedBotBarracks    Barracks `json:"bad_rax_range_bot"`
		RangedTopBarracks    Barracks `json:"bad_rax_range_top"`
		MeleeMiddleBarracks  Barracks `json:"bad_rax_melee_mid"`
		MeleedBotBarracks    Barracks `json:"bad_rax_melee_bot"`
	} `json:"dire"`
}

type Player struct {
	SteamID        string `json:"steamid"`
	Name           string `json:"name"`
	Kills          int8   `json:"kills"`
	Deaths         int8   `json:"deaths"`
	Assists        int8   `json:"assists"`
	LastHits       int8   `json:"last_hits"`
	KillStreak     int8   `json:"kill_streak"`
	TeamName       string `json:"team_name"`
	GPM            int8   `json:"gpm"`
	XPM            int8   `json:"xpm"`
	NetWorth       int16  `json:"net_worth"`
	HeroDamage     int16  `json:"hero_damage"`
	WardsPurchased int8   `json:"wards_purchased"`
	WardsPlaced    int8   `json:"wards_placed"`
	WardsDestroyed int8   `json:"wards_destroyed"`
	RunesActivated int8   `json:"runes_activated"`
	CampsStacked   int8   `json:"camps_stacked"`
}

type PlayerStats struct {
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

type Item struct {
	Purchaser int8   `json:"purchaser"`
	Passive   bool   `json:"passive"`
	Name      string `json:"name"`
}

type PlayerItems struct {
	Slot1  Item `json:"slot1"`
	Slot2  Item `json:"slot2"`
	Slot3  Item `json:"slot3"`
	Slot4  Item `json:"slot4"`
	Slot5  Item `json:"slot5"`
	Slot6  Item `json:"slot6"`
	Slot7  Item `json:"slot7"`
	Slot8  Item `json:"slot8"`
	Stash0 Item `json:"stash0"`
	Stash1 Item `json:"stash1"`
	Stash2 Item `json:"stash2"`
	Stash3 Item `json:"stash3"`
	Stash4 Item `json:"stash4"`
	Stash5 Item `json:"stash5"`
}
type Items struct {
	Radiant struct {
		Player0 PlayerItems `json:"player0"`
		Player1 PlayerItems `json:"player1"`
		Player2 PlayerItems `json:"player2"`
		Player3 PlayerItems `json:"player3"`
		Player4 PlayerItems `json:"player4"`
	} `json:"team2"`
	Dire struct {
		Player5 PlayerItems `json:"player5"`
		Player6 PlayerItems `json:"player6"`
		Player7 PlayerItems `json:"player7"`
		Player8 PlayerItems `json:"player8"`
		Player9 PlayerItems `json:"player9"`
	} `json:"team3"`
}
