package config

type Repository struct {
	API API `json:"api"`
}

type API struct {
	User User `json:"user"`
}

type User struct {
	URL string `json:"url"`
}
