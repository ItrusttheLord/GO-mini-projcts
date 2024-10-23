package models

import (
	"time"
)

type Expenses struct {
	ID          string       `json:"id"`
	Date        time.Time    `json:"date"`
	Information *Information `json:"information"`
	Price       *Price       `json:"price"`
}

type Information struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type Price struct {
	Amount float64 `json:"amount"`
}
