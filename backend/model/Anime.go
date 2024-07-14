package models

type Anime struct {
	Anime_id string `json:"anime_id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Rating   string `json:"rating"`
	Anitype  string `json:"type"`
	Episodes string `json:"episodes"`
	Members  string `json:"members"`
}
