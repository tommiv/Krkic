package fetch

import (
    "fmt"
    "net/http"
    // "github.com/haochi/blockhash-go"
    log "gopkg.in/Sirupsen/logrus.v0"
    "krkic/model"
)

func GetHashByUrl(url string) (*model.UrlInfo, error) {
    log.Debugf("Got new url %s", url)

    resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

    fmt.Printf("%s", resp.Header.Get("content-type"))

	// defer resp.Body.Close()

    // hash, err := blockhash.Blockhash(response.Body, 16)
    // if err != nil {
    //     log.Error(err)
    // }
    //
    // log.Info(hash.ToHex())

    return &model.UrlInfo{}, nil
}
