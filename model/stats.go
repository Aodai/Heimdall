package model

// Stats holds basic stats about a server
type Stats struct {
	Name          string  `json:"name"`
	Screenshot    string  `json:"screenshot"`
	PCBang        int     `json:"PCBang"`
	MaxLevel      int     `json:"max_lvl"`
	PK            int     `json:"PK"`
	DisablePK     int     `json:"disable_pk_on"`
	UserLimit     int     `json:"user_limit"`
	Uptime        int     `json:"uptime"`
	Players       int     `json:"players"`
	ExpRate       float64 `json:"exp_rate"`
	DropRate      float64 `json:"drop_rate"`
	ChaosRate     float64 `json:"chaos_rate"`
	GoldRate      float64 `json:"gold_rate"`
	PartyDropRate float64 `json:"party_drop_rate"`
	PartyExpRate  float64 `json:"party_exp_rate"`
}
