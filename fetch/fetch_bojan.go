package fetch

import (
    "time"
    "strconv"
    "net"
    // "net/url"
    "net/http"
    "strings"
    "krkic/model"
    "krkic/fetch/handlers"
    "github.com/haochi/blockhash-go"
    // log "gopkg.in/Sirupsen/logrus.v0"
)

const DEFAULT_UA = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36"

// TODO: split up semantically
// TODO: make a mutators/handlers list
// TODO: store images on disk as well
// TODO: introduce cache with normalized url as a key
func FetchBojan(job model.FetcherJob) (model.Bojan, error) {
    attempt := model.Attempt {
        BojanedAt: WeirdUnixTimestampToTime(job.Timestamp),
        UserID:    job.OwnerID,
        ChannelID: job.ChannelID,
    }

    answer := model.Bojan {
        URL:      job.URL,
        Type:     model.URLTYPE_OTHER,
        Attempts: []model.Attempt { attempt },
    }

    handler := selectHandler(job)

    if !handler.NeedResponse(job) {
        return answer, nil
    }

    client  := buildClient()
    request, _ := http.NewRequest("GET", job.URL.String(), nil)
    request.Header.Add("User-Agent", DEFAULT_UA)
    request.Header.Add("Accept", "*/*")

    resp, err := client.Do(request)
	if err != nil {
		return answer, err
	}

    defer resp.Body.Close()

    mime := resp.Header.Get("content-type")
    if strings.Contains(mime, "image/") {
        answer.Type = model.URLTYPE_IMAGE

        hash, err := blockhash.Blockhash(resp.Body, 16)
        if err != nil {
            return answer, err
        }

        answer.HashBits = hash.Bits
        answer.HashStr  = hash.ToHex()
    }

    return answer, nil
}

func selectHandler(job model.FetcherJob) handlers.IHandler {
    for _, handler := range handlers.Impl() {
        if handler.CanHandleThis(job) {
            return handler
        }
    }

    basic := handlers.Basic{}
    return  basic // will never get there though
}

// TODO: Configurable timings
func buildClient() *http.Client {
    dialer := &net.Dialer {
      Timeout: 5 * time.Second,
    }

    transport := &http.Transport {
      Dial: dialer.Dial,
      TLSHandshakeTimeout: 5 * time.Second,
    }

    client := &http.Client {
        Timeout:   15 * time.Second,
        Transport: transport,
    }

    return client
}

func WeirdUnixTimestampToTime(input string) time.Time {
    strval := strings.Split(input, ".")[0]
    intval, _ := strconv.Atoi(strval)
    realval := time.Unix(int64(intval), 0)
    return realval
}
