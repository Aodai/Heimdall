package rconsole

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/aodai/heimdall/config"
	"github.com/aodai/heimdall/model"
	"github.com/aodai/heimdall/util"
)

// RConsole blah
type RConsole struct {
	conn   net.Conn
	buffer *bufio.ReadWriter
}

// Init initializes shit
func (rc *RConsole) Init() {
	cfg := config.GetConfig()
	ip := cfg.GetString("Server.IP")
	port := cfg.GetInt("Server.Port")
	pass := cfg.GetString("Server.Password")
	ipEP := fmt.Sprintf("%s:%d", ip, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ipEP)
	if err != nil {
		panic(err)
	}
	rc.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	rc.buffer = bufio.NewReadWriter(bufio.NewReader(rc.conn), bufio.NewWriter(rc.conn))
	_, _, err = rc.buffer.ReadLine()
	if err != nil {
		panic(err)
	}
	_, err = rc.buffer.WriteString(pass + "\n")
	if err != nil {
		panic(err)
	}
	err = rc.buffer.Flush()
	if err != nil {
		panic(err)
	}
	_, _, err = rc.buffer.ReadLine()
	if err != nil {
		panic(err)
	}
	err = rc.buffer.Flush()
	if err != nil {
		panic(err)
	}
}

// Close terminates the connection
func (rc *RConsole) Close() {
	rc.conn.Close()
}

// GetProperty returns the value of a property from Captain Herlock
func (rc *RConsole) GetProperty(prop string) string {
	req := fmt.Sprintf("get %s", prop)
	var line []byte
	_, err := rc.buffer.WriteString(req + "\n")
	if err != nil {
		panic(err)
	}
	rc.buffer.Flush()
	for strings.HasPrefix(string(line), prop) == false {
		rc.buffer.Flush()
		line, _, _ = rc.buffer.ReadLine()
	}
	return string(line[len(prop)+3:])
}

// FetchStats fetches basic stats about the server.
func (rc *RConsole) FetchStats() model.Stats {
	var stats model.Stats
	stats.Name = rc.GetProperty("auth.server_name")
	stats.Screenshot = rc.GetProperty("game.server_screenshot_url")
	stats.PCBang, _ = strconv.Atoi(rc.GetProperty("game.pcbang_bonus_server"))
	stats.MaxLevel, _ = strconv.Atoi(rc.GetProperty("game.max_level"))
	stats.PK, _ = strconv.Atoi(rc.GetProperty("game.PKServer"))
	stats.DisablePK, _ = strconv.Atoi(rc.GetProperty("game.disable_pk_on"))
	stats.UserLimit, _ = strconv.Atoi(rc.GetProperty("set_user_limit"))
	stats.Players, _ = strconv.Atoi(rc.GetProperty("game.user_count"))
	stats.Uptime = util.ConvertToSeconds(rc.GetProperty("process.uptime"))
	stats.ChaosRate, _ = strconv.ParseFloat(rc.GetProperty("game.chaos_drop_rate"), 32)
	stats.DropRate, _ = strconv.ParseFloat(rc.GetProperty("game.item_drop_rate"), 32)
	stats.ExpRate, _ = strconv.ParseFloat(rc.GetProperty("game.exp_rate"), 32)
	stats.GoldRate, _ = strconv.ParseFloat(rc.GetProperty("game.gold_drop_rate"), 32)
	stats.PartyDropRate, _ = strconv.ParseFloat(rc.GetProperty("game.party_drop_rate"), 32)
	stats.PartyExpRate, _ = strconv.ParseFloat(rc.GetProperty("game.party_exp_rate"), 32)
	return stats
}
