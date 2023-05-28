package models

type Names struct {
	Names string `json:"names"`
}

type SoundexResponse struct {
	Soundexes []string `json:"soundex"`
}
