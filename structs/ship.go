package structs

// Names the name of the ship in different languages
type Names struct {
	En string
	Cn string
	Jp string
	Kr string
}

// Skin a specific ships skin
type Skin struct {
	Title string
	Image string
	Chibi string
}

// Stars how many stars the ship starts with
type Stars struct {
	Value string
	Count int
}

// Stat represents a single stat object
type Stat struct {
	Name  string
	Image string
	Value string
}

// Stats lists of the ships stats
type Stats struct {
	Level100    []Stat
	Level120    []Stat
	Base        []Stat
	Retrofit100 []Stat
	Retrofit120 []Stat
}

// MiscellaneousData link and name of artists
type MiscellaneousData struct {
	Link string
	Name string
}

// Miscellaneous links and names of artists that have worked on the ship
type Miscellaneous struct {
	Artist       MiscellaneousData
	Web          MiscellaneousData
	Pixiv        MiscellaneousData
	Twitter      MiscellaneousData
	VoiceActress MiscellaneousData
}

// SmallShip ship but only the id and name
type SmallShip struct {
	ID   string
	Name string
}

// Ship represents an azur lane ship
type Ship struct {
	WikiURL          string
	ID               string
	Name             string
	Thumbnail        string
	Skins            []Skin
	BuildTime        string
	Rarity           string
	Stars            Stars
	Class            string
	Nationality      string
	NationalityShort string
	HullType         string
	Stats            Stats
	Miscellaneous    Miscellaneous
}
