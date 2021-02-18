package csmmapi

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type (
	//GameServer ....
	GameServer struct {
		IP         string
		Port       string
		AdminUser  string
		AdminToken string
		scheme     string
		client     *resty.Client
	}
	//GameTime ....
	GameTime struct {
		Days    json.Number `json:"days"`
		Hours   json.Number `json:"hours"`
		Minutes json.Number `json:"minutes"`
	}
)

//Stats ....
type Stats struct {
	Gametime GameTime    `json:"gametime"`
	Players  json.Number `json:"players"`
	Hostiles json.Number `json:"hostiles"`
	Animals  json.Number `json:"animals"`
}

//Position ....
type Position struct {
	X json.Number `json:"x"`
	Y json.Number `json:"y"`
	Z json.Number `json:"z"`
}

//OnlinePlayer ....
type OnlinePlayer struct {
	SteamID      string      `json:"steamid"`
	EntityID     json.Number `json:"entityid"`
	IP           string      `json:"ip"`
	Name         string      `json:"name"`
	Online       bool        `json:"online"`
	Position     Position    `json:"position"`
	Experience   json.Number `json:"experience"`
	Level        json.Number `json:"level"`
	Health       json.Number `json:"health"`
	Stamina      json.Number `json:"stamina"`
	ZombieKills  json.Number `json:"zombiekills"`
	PlayerKills  json.Number `json:"playerkills"`
	PlayerDeaths json.Number `json:"deaths"`
	Score        json.Number `json:"score"`
	Playtime     json.Number `json:"totalplaytime"`
	LastOnline   string      `json:"lastonline"`
	Ping         json.Number `json:"ping"`
}

//OnlinePlayers ....
type OnlinePlayers []OnlinePlayer

//AllowedCommands ....
type AllowedCommands struct {
	Commands []Command `json:"commands"`
}

//Command ....
type Command struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	Help        string `json:"help"`
}

//CommandResponse ...
type CommandResponse struct {
	Command    string `json:"command"`
	Parameters string `json:"parameters"`
	Result     string `json:"result"`
}

//EntityLocations ....
type EntityLocations []EntityLocationElement

//EntityLocationElement ....
type EntityLocationElement struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

//LandClaims ....
type LandClaims struct {
	Claimsize   int64        `json:"claimsize"`
	ClaimOwners []ClaimOwner `json:"claimowners"`
}

//ClaimOwner ....
type ClaimOwner struct {
	Steamid     string     `json:"steamid"`
	Claimactive bool       `json:"claimactive"`
	Playername  string     `json:"playername"`
	Claims      []Position `json:"claims"`
}

//Inventory ....
type Inventory struct {
	Steamid    string    `json:"steamid"`
	Entityid   int64     `json:"entityid"`
	Playername string    `json:"playername"`
	Bag        []*Belt   `json:"bag"`
	Belt       []Belt    `json:"belt"`
	Equipment  Equipment `json:"equipment"`
}

//Belt ....
type Belt struct {
	Count        int64   `json:"count"`
	Name         string  `json:"name"`
	Icon         string  `json:"icon"`
	Iconcolor    string  `json:"iconcolor"`
	Quality      int64   `json:"quality"`
	Qualitycolor *string `json:"qualitycolor,omitempty"`
}

//Equipment ....
type Equipment struct {
	Head     Belt `json:"head"`
	Eyes     Belt `json:"eyes"`
	Face     Belt `json:"face"`
	Armor    Belt `json:"armor"`
	Jacket   Belt `json:"jacket"`
	Shirt    Belt `json:"shirt"`
	Legarmor Belt `json:"legarmor"`
	Pants    Belt `json:"pants"`
	Boots    Belt `json:"boots"`
	Gloves   Belt `json:"gloves"`
}

//Inventories ....
type Inventories []Inventory

//PlayerList ....
type PlayerList struct {
	Total           int64    `json:"total"`
	TotalUnfiltered int64    `json:"totalUnfiltered"`
	FirstResult     int64    `json:"firstResult"`
	Players         []Player `json:"players"`
}

//Player ....
type Player struct {
	Steamid       string   `json:"steamid"`
	Entityid      int64    `json:"entityid"`
	IP            string   `json:"ip"`
	Name          string   `json:"name"`
	Online        bool     `json:"online"`
	Position      Position `json:"position"`
	Totalplaytime int64    `json:"totalplaytime"`
	Lastonline    string   `json:"lastonline"`
	Ping          int64    `json:"ping"`
	Banned        bool     `json:"banned"`
}

//PlayerLocations ....
type PlayerLocations []PlayerLocation

//PlayerLocation ....
type PlayerLocation struct {
	Steamid  string   `json:"steamid"`
	Name     string   `json:"name"`
	Online   bool     `json:"online"`
	Position Position `json:"position"`
}

//ServerInfo ....
type ServerInfo struct {
	GameType                           StringType `json:"GameType"`
	GameName                           StringType `json:"GameName"`
	GameHost                           StringType `json:"GameHost"`
	ServerDescription                  StringType `json:"ServerDescription"`
	ServerWebsiteURL                   StringType `json:"ServerWebsiteURL"`
	LevelName                          StringType `json:"LevelName"`
	GameMode                           StringType `json:"GameMode"`
	Version                            StringType `json:"Version"`
	IP                                 StringType `json:"IP"`
	CountryCode                        StringType `json:"CountryCode"`
	SteamID                            StringType `json:"SteamID"`
	CompatibilityVersion               StringType `json:"CompatibilityVersion"`
	Platform                           StringType `json:"Platform"`
	ServerLoginConfirmationText        StringType `json:"ServerLoginConfirmationText"`
	Port                               IntType    `json:"Port"`
	CurrentPlayers                     IntType    `json:"CurrentPlayers"`
	MaxPlayers                         IntType    `json:"MaxPlayers"`
	GameDifficulty                     IntType    `json:"GameDifficulty"`
	DayNightLength                     IntType    `json:"DayNightLength"`
	BloodMoonFrequency                 IntType    `json:"BloodMoonFrequency"`
	BloodMoonRange                     IntType    `json:"BloodMoonRange"`
	BloodMoonWarning                   IntType    `json:"BloodMoonWarning"`
	ZombiesRun                         IntType    `json:"ZombiesRun"`
	ZombieMove                         IntType    `json:"ZombieMove"`
	ZombieMoveNight                    IntType    `json:"ZombieMoveNight"`
	ZombieFeralMove                    IntType    `json:"ZombieFeralMove"`
	ZombieBMMove                       IntType    `json:"ZombieBMMove"`
	XPMultiplier                       IntType    `json:"XPMultiplier"`
	DayCount                           IntType    `json:"DayCount"`
	Ping                               IntType    `json:"Ping"`
	DropOnDeath                        IntType    `json:"DropOnDeath"`
	DropOnQuit                         IntType    `json:"DropOnQuit"`
	BloodMoonEnemyCount                IntType    `json:"BloodMoonEnemyCount"`
	EnemyDifficulty                    IntType    `json:"EnemyDifficulty"`
	PlayerKillingMode                  IntType    `json:"PlayerKillingMode"`
	CurrentServerTime                  IntType    `json:"CurrentServerTime"`
	DayLightLength                     IntType    `json:"DayLightLength"`
	BlockDurabilityModifier            IntType    `json:"BlockDurabilityModifier"`
	BlockDamagePlayer                  IntType    `json:"BlockDamagePlayer"`
	BlockDamageAI                      IntType    `json:"BlockDamageAI"`
	BlockDamageAIBM                    IntType    `json:"BlockDamageAIBM"`
	AirDropFrequency                   IntType    `json:"AirDropFrequency"`
	LootAbundance                      IntType    `json:"LootAbundance"`
	LootRespawnDays                    IntType    `json:"LootRespawnDays"`
	MaxSpawnedZombies                  IntType    `json:"MaxSpawnedZombies"`
	LandClaimCount                     IntType    `json:"LandClaimCount"`
	LandClaimSize                      IntType    `json:"LandClaimSize"`
	LandClaimDeadZone                  IntType    `json:"LandClaimDeadZone"`
	LandClaimExpiryTime                IntType    `json:"LandClaimExpiryTime"`
	LandClaimDecayMode                 IntType    `json:"LandClaimDecayMode"`
	LandClaimOnlineDurabilityModifier  IntType    `json:"LandClaimOnlineDurabilityModifier"`
	LandClaimOfflineDurabilityModifier IntType    `json:"LandClaimOfflineDurabilityModifier"`
	LandClaimOfflineDelay              IntType    `json:"LandClaimOfflineDelay"`
	PartySharedKillRange               IntType    `json:"PartySharedKillRange"`
	MaxSpawnedAnimals                  IntType    `json:"MaxSpawnedAnimals"`
	ServerVisibility                   IntType    `json:"ServerVisibility"`
	BedrollExpiryTime                  IntType    `json:"BedrollExpiryTime"`
	IsDedicated                        BoolType   `json:"IsDedicated"`
	IsPasswordProtected                BoolType   `json:"IsPasswordProtected"`
	ShowFriendPlayerOnMap              BoolType   `json:"ShowFriendPlayerOnMap"`
	BuildCreate                        BoolType   `json:"BuildCreate"`
	EACEnabled                         BoolType   `json:"EACEnabled"`
	Architecture64                     BoolType   `json:"Architecture64"`
	StockSettings                      BoolType   `json:"StockSettings"`
	StockFiles                         BoolType   `json:"StockFiles"`
	ModdedConfig                       BoolType   `json:"ModdedConfig"`
	RequiresMod                        BoolType   `json:"RequiresMod"`
	AirDropMarker                      BoolType   `json:"AirDropMarker"`
	EnemySpawnMode                     BoolType   `json:"EnemySpawnMode"`
	IsPublic                           BoolType   `json:"IsPublic"`
}

//IntType ....
type IntType struct {
	Type  string      `json:"type"`
	Value json.Number `json:"value"`
}

//BoolType ....
type BoolType struct {
	Type  string `json:"type"`
	Value bool   `json:"value"`
}

//StringType ....
type StringType struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//WebUIUpdates ....
type WebUIUpdates struct {
	Gametime GameTime `json:"gametime"`
	Players  int64    `json:"players"`
	Hostiles int64    `json:"hostiles"`
	Animals  int64    `json:"animals"`
	Newlogs  int64    `json:"newlogs"`
}

//Log ....
type Log struct {
	FirstLine int64      `json:"firstLine"`
	LastLine  int64      `json:"lastLine"`
	Entries   []LogEntry `json:"entries"`
}

//LogEntry ....
type LogEntry struct {
	Date   string `json:"date"`
	Time   string `json:"time"`
	Uptime string `json:"uptime"`
	Msg    string `json:"msg"`
	Trace  string `json:"trace"`
	Type   string `json:"type"`
}
