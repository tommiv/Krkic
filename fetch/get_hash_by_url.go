package fetch

import (
    "time"
    "net"
    "net/url"
    "net/http"
    "strings"
    "krkic/model"
    "github.com/haochi/blockhash-go"
    // log "gopkg.in/Sirupsen/logrus.v0"
)

const DEFAULT_UA = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36"

// TODO: split up semantically
// TODO: rename to what it really does
// TODO: make a mutators/hanlders list
// TODO: store images on disk as well
// TODO: introduce cache with normalized url as a key
func GetHashByUrl(url *url.URL) (*model.Bojan, error) {
    answer := &model.Bojan {
        URL:  url,
        Type: model.URLTYPE_OTHER,
    }

    client  := buildClient()
    request, _ := http.NewRequest("GET", url.String(), nil)
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

// TODO: Configurable timings
func buildClient() *http.Client {
    dialer := &net.Dialer {
      Timeout: 2 * time.Second,
    }

    transport := &http.Transport {
      Dial: dialer.Dial,
      TLSHandshakeTimeout: 2 * time.Second,
    }

    client := &http.Client {
        Timeout:   5 * time.Second,
        Transport: transport,
    }

    return client
}
