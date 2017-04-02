package model

import (
    "time"
)

type Attempt struct {
    BojanedAt time.Time
    UserID    string
    ChannelID string
}
