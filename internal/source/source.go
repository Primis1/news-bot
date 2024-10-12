package source

import (
	"context"
	"github.com/SlyMarbo/rss"
	"github.com/samber/lo"
	"strings"

	model "news-bot/internal/model"
)

type RSSSource struct {
	URL        string
	SourceID   int64
	SourceName string
}

func NewRSSSourceFromModel(m model.Source) RSSSource {
	return RSSSource{
		URL:        m.FeedURL,
		SourceID:   m.ID,
		SourceName: m.Name,
	}
}

func (s RSSSource) Fetch(ctx context.Context) ([]model.Item, error) {
	feed, err := s.LoadSource(ctx, s.URL)
	if err != nil {
		return nil, err
	}

	return lo.Map(feed.Items, func(item *rss.Item, _ int) model.Item {
		return model.Item{
			Title:      item.Title,
			Categories: item.Categories,
			Link:       item.Link,
			Date:       item.Date,
			SourceName: s.SourceName,
			Summary:    strings.TrimSpace(item.Summary)}
	}), nil

}

func (s RSSSource) LoadSource(ctx context.Context, url string) (r *rss.Feed, err error) {
	var (
		feedCh  = make(chan *rss.Feed)
		errorCh = make(chan error)
	)

	go func() {
		feed, err := rss.Fetch(url)
		if err != nil {
			errorCh <- err
			return
		}

		feedCh <- feed
	}()

	select {
	// case of finishing context
	case <-ctx.Done():
		return nil, ctx.Err()

		// case of error in context
	case err := <-errorCh:
		return nil, err
		// case of successful, return feed as an object
	case feed := <-feedCh:
		return feed, nil
	}
}
