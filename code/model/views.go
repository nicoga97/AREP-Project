package model

type Store struct {
	Id        int64      `json:"id"`
	StoreId   int64      `json:"store_id"`
	BName     string     `json:"b_name"`
	Schedules []Schedule `json:"schedules"`
	Products  []Product  `json:"products"`
	Name      string     `json:"name"`
	Enabled   bool       `json:"enabled"`
}

type Schedule struct {
	Id    int64  `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Product struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}
