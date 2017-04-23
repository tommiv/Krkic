package fetch

import (
	_ "image/gif" // todo
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
    "io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
    "krkic/di"
    "krkic/fetch/handlers"
    "krkic/model"
    "github.com/haochi/blockhash-go"
	"gopkg.in/spf13/viper.v0"
	"gopkg.in/vmihailenco/msgpack.v2"
	// log "gopkg.in/Sirupsen/logrus.v0"
)

const defaultUA = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36"

// FetchBojan Fetch bojan from job definition
// TODO: split up semantically
// TODO: make a mutators/handlers list
func FetchBojan(job model.FetcherJob) (model.Bojan, error) {
	attempt := model.Attempt{
		BojanedAt: weirdUnixTimestampToTime(job.Timestamp),
		UserID:    job.OwnerID,
		ChannelID: job.ChannelID,
	}

	handler := selectHandler(job)

	url := handler.ModifyURL(job.URL)
	urlhash := urlToMD5Hex(url)

	cacheKey := "cache:http:" + urlhash

	answer := model.Bojan{
		URL:      url,
		Type:     model.URLTYPE_OTHER,
		Attempts: []model.Attempt{attempt},
	}

	if !handler.NeedResponse(job) {
		return answer, nil
	}

	storedBuf, _ := di.RedisOther.Get(cacheKey).Bytes()
	if len(storedBuf) > 0 {
		var stored model.Bojan
		err := msgpack.Unmarshal(storedBuf, &stored)

		if err == nil {
			return stored, nil
		}
	}

	client := buildClient()
	request, _ := http.NewRequest("GET", url.String(), nil)
	request.Header.Add("User-Agent", defaultUA)
	request.Header.Add("Accept", "*/*")

	resp, err := client.Do(request)
	if err != nil {
		return answer, err
	}

	defer resp.Body.Close()

	mime := resp.Header.Get("content-type")
	mimeChunks := strings.Split(mime, "/")

	if len(mimeChunks) == 2 && mimeChunks[0] == "image" {
		answer.Type = model.URLTYPE_IMAGE

		buf, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return answer, err
        }

        stream := bytes.NewReader(buf)

		hash, err := blockhash.Blockhash(stream, 16)
		if err != nil {
            msg := err.Error() + " (" + mimeChunks[1] + ")"
			return answer, errors.New(msg)
		}

		answer.HashBits = hash.Bits
		answer.HashStr = hash.ToHex()

        fileName := urlhash + "." + mimeChunks[1]
        fileDumpPath := path.Join(viper.GetString("app_data_folder"), "imgcache", fileName)
        // ignoring errors for now
        ioutil.WriteFile(fileDumpPath, buf, 0777)
	}

	tostore, _ := msgpack.Marshal(answer)
	di.RedisOther.Set(cacheKey, tostore, 0)

	return answer, nil
}

func selectHandler(job model.FetcherJob) handlers.IHandler {
	for _, handler := range handlers.Impl() {
		if handler.CanHandleThis(job) {
			return handler
		}
	}

	basic := handlers.Basic{}
	return basic // will never get there though
}

// TODO: Configurable timings
func buildClient() *http.Client {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	transport := &http.Transport{
		Dial:                dialer.Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: transport,
	}

	return client
}

func weirdUnixTimestampToTime(input string) time.Time {
	strval := strings.Split(input, ".")[0]
	intval, _ := strconv.Atoi(strval)
	realval := time.Unix(int64(intval), 0)
	return realval
}

func urlToMD5Hex(url url.URL) string {
	inputBuf := []byte(url.String())
	sum := md5.Sum(inputBuf)
	return hex.EncodeToString(sum[:])
}
