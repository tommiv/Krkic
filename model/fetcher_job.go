package model

import (
    "net/url"
)

type FetcherJob struct {
    URL       *url.URL
    Title     string // TODO: why
    MimeType  string
    OwnerID   string
    ChannelID string
    Timestamp string
}
