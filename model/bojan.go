package model

import (
    "net/url"
)

const URLTYPE_IMAGE = 0
const URLTYPE_OTHER = 1

type Bojan struct {
    URL      url.URL
    Type     int
    HashBits []int
    HashStr  string
    Attempts []Attempt
}
