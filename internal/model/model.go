package model

import "time"

type Item struct {
	Title       string
	Link        string
	Categories  []string
	PublishedAt string
	Summary     string
	Date        time.Time
}

type Source struct {
	ID        int64
	Name      string
	Feed      string
	CraetedAt time.Time
}

type Article struct {
	ID        int64
	SourceID  int64
	Title     string
	Link      string
	Summary   string
	Published time.Time
	PostedAt  time.Time
	CreateAt  time.Time
}
