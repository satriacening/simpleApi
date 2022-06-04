package entities

type Response struct {
	Type    string          `json:"type"`
	Format  string          `json:"format"`
	Version string          `json:"version"`
	Data    map[string]Data `json:"data"`
}
type Data struct {
	Version string             `json:"version"`
	Id      string             `json:"id"`
	Key     string             `json:"key"`
	Name    string             `json:"Ahri"`
	Title   string             `json:"title"`
	Blurb   string             `json:"blurb"`
	Info    Info               `json:"info"`
	Image   Image              `json:"image"`
	Tags    []Tag              `json:"tags"`
	Partype Partype            `json:"partype"`
	Stats   map[string]float64 `json:"stats"`
}

type Image struct {
	Full   string `json:"full"`
	Sprite Sprite `json:"sprite"`
	Group  Type   `json:"group"`
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
	W      int64  `json:"w"`
	H      int64  `json:"h"`
}

type Info struct {
	Attack     int64 `json:"attack"`
	Defense    int64 `json:"defense"`
	Magic      int64 `json:"magic"`
	Difficulty int64 `json:"difficulty"`
}

type Type string

const (
	Champion Type = "champion"
)

type Sprite string

const (
	Champion0PNG Sprite = "champion0.png"
	Champion1PNG Sprite = "champion1.png"
	Champion2PNG Sprite = "champion2.png"
	Champion3PNG Sprite = "champion3.png"
	Champion4PNG Sprite = "champion4.png"
)

type Partype string

const (
	Battlefury Partype = "Battlefury"
	BloodWell  Partype = "BloodWell"
	Dragonfury Partype = "Dragonfury"
	Energy     Partype = "Energy"
	Ferocity   Partype = "Ferocity"
	Gnarfury   Partype = "Gnarfury"
	Heat       Partype = "Heat"
	Mp         Partype = "MP"
	None       Partype = "None"
	Rage       Partype = "Rage"
	Wind       Partype = "Wind"
)

type Tag string

const (
	Assassin Tag = "Assassin"
	Fighter  Tag = "Fighter"
	Mage     Tag = "Mage"
	Marksman Tag = "Marksman"
	Melee    Tag = "Melee"
	Support  Tag = "Support"
	Tank     Tag = "Tank"
)

type Version string

const (
	The6241 Version = "6.24.1"
)
