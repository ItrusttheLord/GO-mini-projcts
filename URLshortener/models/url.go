package models

import (
	"time"
)

type URL struct {
	ShortURL    string `json:"shortUrl"`
	OriginalURL string `json:"originalUrl"`
	Date        time.Time
}
