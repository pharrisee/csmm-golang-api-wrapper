package csmmapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

//NewGameServer ....
func NewGameServer(ip string, port string, user string, token string, schemeIn ...string) *GameServer {
	scheme := "http"
	if len(schemeIn) > 0 {
		scheme = schemeIn[0]
	}
	return &GameServer{
		IP:         ip,
		Port:       port,
		AdminUser:  user,
		AdminToken: token,
		scheme:     scheme,
		client:     resty.New(),
	}
}

func (gs GameServer) baseURL() url.URL {
	return url.URL{
		Scheme: gs.scheme,
		Host:   fmt.Sprintf("%s:%s", gs.IP, gs.Port),
	}
}

func (gs GameServer) fetchJSON(method string, retval interface{}, queries ...url.Values) (err error) {
	query := url.Values{}
	if len(queries) != 0 {
		query = queries[0]
	}
	uri := gs.baseURL()
	uri.Path = method
	query.Add("adminuser", gs.AdminUser)
	query.Add("admintoken", gs.AdminToken)
	uri.RawQuery = query.Encode()

	fmt.Println(uri.String())
	resp, err := gs.client.R().Get(uri.String())
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("%d: %s", resp.StatusCode(), resp.Status())
	}
	err = json.Unmarshal(resp.Body(), retval)
	return
}

//GetStats ....
func (gs GameServer) GetStats() (stats Stats, err error) {
	err = gs.fetchJSON("/api/getstats", &stats)
	return
}

//GetOnlinePlayers ....
func (gs GameServer) GetOnlinePlayers() (onlinePlayers OnlinePlayers, err error) {
	err = gs.fetchJSON("/api/getplayersonline", &onlinePlayers)
	return
}

//GetAllowedCommands ....
func (gs GameServer) GetAllowedCommands() (allowedCommands AllowedCommands, err error) {
	err = gs.fetchJSON("/api/getallowedcommands", &allowedCommands)
	return
}

//ExecuteConsoleCommand ....
func (gs GameServer) ExecuteConsoleCommand(command string) (commandResponse CommandResponse, err error) {
	params := url.Values{}
	params.Add("command", command)
	err = gs.fetchJSON("/api/executeconsolecommand", &commandResponse, params)
	return
}

//GetAnimalsLocation ....
func (gs GameServer) GetAnimalsLocation() (entityLocations EntityLocations, err error) {
	err = gs.fetchJSON("/api/getanimalslocation", &entityLocations)
	return
}

//GetHostileLocation ....
func (gs GameServer) GetHostileLocation() (entityLocations EntityLocations, err error) {
	err = gs.fetchJSON("/api/gethostilelocation", &entityLocations)
	return
}

//GetLandClaims ....
func (gs GameServer) GetLandClaims(steamID string) (landClaims LandClaims, err error) {
	params := url.Values{}
	params.Add("steamid", steamID)
	err = gs.fetchJSON("/api/getlandclaims", &landClaims, params)
	return
}

//GetPlayerInventory ....
func (gs GameServer) GetPlayerInventory(steamID string) (inventory Inventory, err error) {
	params := url.Values{}
	params.Add("steamid", steamID)
	err = gs.fetchJSON("/api/getplayerinventory", &inventory, params)
	return
}

//GetPlayerInventories ....
func (gs GameServer) GetPlayerInventories() (inventories Inventories, err error) {
	err = gs.fetchJSON("/api/getplayerinventories", &inventories)
	return
}

//GetPlayerList ....
func (gs GameServer) GetPlayerList(rowsperpage int, page int) (playerList PlayerList, err error) {
	params := url.Values{}
	params.Add("rowsperpage", fmt.Sprintf("%d", rowsperpage))
	params.Add("page", fmt.Sprintf("%d", page))
	err = gs.fetchJSON("/api/getplayerlist", &playerList, params)
	return
}

//GetPlayersLocation ....
func (gs GameServer) GetPlayersLocation(offline bool) (playerLocations PlayerLocations, err error) {
	params := url.Values{}
	params.Add("offline", fmt.Sprintf("%t", offline))
	err = gs.fetchJSON("/api/getplayerslocation", &playerLocations, params)
	return
}

//GetServerInfo ....
func (gs GameServer) GetServerInfo() (serverInfo ServerInfo, err error) {
	err = gs.fetchJSON("/api/getserverinfo", &serverInfo)
	return
}

//GetWebUIUpdates ....
func (gs GameServer) GetWebUIUpdates(latestLine int) (webUIUpdates WebUIUpdates, err error) {
	params := url.Values{}
	params.Add("latestline", fmt.Sprintf("%d", latestLine))
	err = gs.fetchJSON("/api/getwebuiupdates", &webUIUpdates, params)
	return
}

//GetLog ....
func (gs GameServer) GetLog(firstLine int, counts ...int) (theLog Log, err error) {
	count := 50
	if len(counts) > 0 {
		count = counts[0]
	}
	firstLineS := fmt.Sprintf("%d", firstLine)
	params := url.Values{}
	params.Add("firstline", firstLineS)
	params.Add("count", fmt.Sprintf("%d", count))
	err = gs.fetchJSON("/api/getlog", &theLog, params)
	return
}
