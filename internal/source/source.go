package source

import (
	"context"
	"github.com/SlyMarbo/rss"
	"github.com/samber/lo"
	"news-bot/internal/model")

type RSSSource struct {
	URL        string
	SourceID   int64
	SourceName string
}

func NameRSSSource(m model.Source) RSSSource{
	return RSSSource{
		URL: m.Feed,
		SourceID: m.ID,
		SourceName: m.Name,
	}
}