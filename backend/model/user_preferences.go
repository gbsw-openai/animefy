package models

type UserPreferences struct {
	Genres      []string `json:"genres"`
	MinRating   float64  `json:"min_rating"`
	MaxEpisodes int      `json:"max_episodes"`
	Type        string   `json:"type"`
}
