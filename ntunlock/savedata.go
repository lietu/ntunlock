package ntunlock

type Options struct {
	Ambvol       float32 `json:"ambvol"`
	Autoaim      float32 `json:"autoaim"`
	Bintro       float32 `json:"bintro"`
	Borderless   float32 `json:"borderless"`
	Coop         float32 `json:"coop"`
	Crosshair    float32 `json:"crosshair"`
	Fitscrn      float32 `json:"fitscrn"`
	Freeze       float32 `json:"freeze"`
	Fulscrn      float32 `json:"fulscrn"`
	Gamepad      float32 `json:"gamepad"`
	Gamepadstyle float32 `json:"gamepadstyle"`
	Gpkeyfire    float32 `json:"gpkeyfire"`
	Gpkeypick    float32 `json:"gpkeypick"`
	Gpkeyspec    float32 `json:"gpkeyspec"`
	Gpkeyswap    float32 `json:"gpkeyswap"`
	Kbkeydown    float32 `json:"kbkeydown"`
	Kbkeyfire    float32 `json:"kbkeyfire"`
	Kbkeyleft    float32 `json:"kbkeyleft"`
	Kbkeypick    float32 `json:"kbkeypick"`
	Kbkeyright   float32 `json:"kbkeyright"`
	Kbkeyspec    float32 `json:"kbkeyspec"`
	Kbkeyswap    float32 `json:"kbkeyswap"`
	Kbkeyup      float32 `json:"kbkeyup"`
	Lfthand      float32 `json:"lfthand"`
	Mousecp      float32 `json:"mousecp"`
	Musvol       float32 `json:"musvol"`
	Nicedrk      float32 `json:"nicedrk"`
	Niceshd      float32 `json:"niceshd"`
	Nohud        float32 `json:"nohud"`
	Rumble       float32 `json:"rumble"`
	Scaling      float32 `json:"scaling"`
	Sfxvol       float32 `json:"sfxvol"`
	Shake        float32 `json:"shake"`
	Sideart      float32 `json:"sideart"`
	Streamkey    string  `json:"streamkey"`
	Timer        float32 `json:"timer"`
}

type Character struct {
	Crowns     []float32 `json:"crowns"`
	Hbst_area  float32   `json:"hbst_area"`
	Ctot_loop  float32   `json:"ctot_loop"`
	Hbst_loop  float32   `json:"hbst_loop"`
	Cbst_strk  float32   `json:"cbst_strk"`
	Cbst_suba  float32   `json:"cbst_suba"`
	Startskin  float32   `json:"startskin"`
	Ctot_days  float32   `json:"ctot_days"`
	Ctot_hard  float32   `json:"ctot_hard"`
	Dbst_area  float32   `json:"dbst_area"`
	Dbst_loop  float32   `json:"dbst_loop"`
	Cgld       float32   `json:"cgld"`
	Ctot_dead  float32   `json:"ctot_dead"`
	Ctot_uniq  float32   `json:"ctot_uniq"`
	Ctot_time  float32   `json:"ctot_time"`
	Dbst_race  float32   `json:"dbst_race"`
	Cbst_kill  float32   `json:"cbst_kill"`
	Ctot_runs  float32   `json:"ctot_runs"`
	Hbst_race  float32   `json:"hbst_race"`
	Cgot       float32   `json:"cgot"`
	Cwep       float32   `json:"cwep"`
	Dbst_kill  float32   `json:"dbst_kill"`
	Cbst_race  float32   `json:"cbst_race"`
	Cbst_fast  float32   `json:"cbst_fast"`
	Hbst_kill  float32   `json:"hbst_kill"`
	Ctot_kill  float32   `json:"ctot_kill"`
	Dbst_suba  float32   `json:"dbst_suba"`
	Ctot_wins  float32   `json:"ctot_wins"`
	Cbst_loop  float32   `json:"cbst_loop"`
	Ctot_strk  float32   `json:"ctot_strk"`
	Cbst_area  float32   `json:"cbst_area"`
	Cbgt       float32   `json:"cbgt"`
	Hbst_suba  float32   `json:"hbst_suba"`
	Startcrown float32   `json:"startcrown"`
}

type Stats struct {
	Fish     *Character `json:"charData_1"`
	Crystal  *Character `json:"charData_2"`
	Eyes     *Character `json:"charData_3"`
	Melting  *Character `json:"charData_4"`
	Plant    *Character `json:"charData_5"`
	YV       *Character `json:"charData_6"`
	Steroids *Character `json:"charData_7"`
	Robot    *Character `json:"charData_8"`
	Chicken  *Character `json:"charData_9"`
	Rebel    *Character `json:"charData_10"`
	Horror   *Character `json:"charData_11"`
	Rogue    *Character `json:"charData_12"`
	BigDog   *Character `json:"charData_13"`
	Skeleton *Character `json:"charData_14"`
	Frog     *Character `json:"charData_15"`
	YC       *Character `json:"charData_16"`
	Tot_time float32    `json:"tot_time"`
}

type Save struct {
	DRM     string                `json:"VLAMBEER DRM 2013-NOW"`
	Options *Options              `json:"options"`
	Stats   *Stats                `json:"stats"`
	Version string                `json:"version"`
	Data    map[string]float32    `json:"data"`
}

func NewCharacter() *Character {
	c := Character{
		[]float32{},
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	}

	return &c
}

func NewStats() *Stats {
	s := Stats{
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		NewCharacter(),
		0.0,
	}

	return &s
}

func NewOptions() *Options {
	o := Options{
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
		"",
		0.0,
	}

	return &o
}

func NewSave() *Save {
	s := Save{
		"",
		NewOptions(),
		NewStats(),
		"",
		map[string]float32{},
	}

	return &s
}
