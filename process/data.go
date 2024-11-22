package main

type Emoji struct {
	Path  string
	Width int
	Alt   string
}

type EmojiRow struct {
	Emojis []Emoji
}

type EmojiInclude struct {
	Name  string              `json:"name"`
	Type  string              `json:"type"`
	Items []map[string]string `json:"items"`
}
