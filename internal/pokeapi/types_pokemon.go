package pokeapi

type PokemonDetails struct {
	ID                     int           `json:"id"`
	Name                   string        `json:"name"`
	BaseExperience         int           `json:"base_experience"`
	Height                 int           `json:"height"`
	IsDefault              bool          `json:"is_default"`
	Order                  int           `json:"order"`
	Weight                 int           `json:"weight"`
	Abilities              []Abilities   `json:"abilities"`
	Forms                  []Forms       `json:"forms"`
	GameIndices            []GameIndices `json:"game_indices"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Moves       `json:"moves"`
	Species                Species       `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Cries                  Cries         `json:"cries"`
	Stats                  []Stats       `json:"stats"`
	Types                  []Types       `json:"types"`
	PastTypes              []PastTypes   `json:"past_types"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	VersionGroup    VersionGroup    `json:"version_group"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
}
type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Sprites struct {
	BackDefault      *string   `json:"back_default"`
	BackFemale       *string      `json:"back_female"`
	BackShiny        *string   `json:"back_shiny"`
	BackShinyFemale  *string      `json:"back_shiny_female"`
	FrontDefault     *string   `json:"front_default"`
	FrontFemale      *string      `json:"front_female"`
	FrontShiny       *string   `json:"front_shiny"`
	FrontShinyFemale *string      `json:"front_shiny_female"`
}
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PastTypes struct {
	Generation Generation `json:"generation"`
	Types      []Types    `json:"types"`
}