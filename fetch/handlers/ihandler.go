package handlers

import (
    "net/url"
    "krkic/model"
)

type IHandler interface {
    CanHandleThis(job model.FetcherJob) bool
    NeedResponse (job model.FetcherJob) bool
    ShouldIgnore (job model.FetcherJob) bool
    ModifyURL    (url url.URL) url.URL
    // HttpClient ModifyClient(HttpClient client)
}
