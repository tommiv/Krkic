package handlers

import (
    "net/url"
    "krkic/model"
)

type Basic struct {}

func (this Basic) CanHandleThis(job model.FetcherJob) bool {
    return true
}

func (this Basic) NeedResponse(job model.FetcherJob) bool {
    return true
}

func (this Basic) ShouldIgnore(job model.FetcherJob) bool {
    return false
}

func (this Basic) ModifyURL(url url.URL) url.URL {
    return url
}
