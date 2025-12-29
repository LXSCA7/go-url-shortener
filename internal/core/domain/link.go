package domain

import "time"

type Link struct {
	Id          int64     `json"id,string"`
	OriginalURL string    `json:"original_url"`
	ShortCOde   string    `json:"short_code"`
	CreatedAt   time.Time `json:"created_at"`
	Visits      int       `json:"visits"`
}
