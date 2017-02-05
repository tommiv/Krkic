package training

import (
    "strings"
	"net/url"
    "krkic/model"
	"github.com/mvdan/xurls"
    log "gopkg.in/Sirupsen/logrus.v0"
)

func ParseJobs(messages []model.ArchivedMessage) []model.FetcherJob {
    var jobs []model.FetcherJob

    for _, msg := range messages {
        found := Message2Jobs(&msg)
        jobs = append(jobs, found...)
    }

    return jobs
}

func Message2Jobs(message *model.ArchivedMessage) []model.FetcherJob {
	if message.Subtype == "" && len(message.Attch) > 0 {
		return attachments2Jobs(message)
    } else if message.Subtype == "" {
        return genericMessage2Jobs(message)
	} else if message.Subtype == "file_share" {
        return fileShares2Jobs(message)
    }

    return make([]model.FetcherJob, 0)
}

func attachments2Jobs(message *model.ArchivedMessage) []model.FetcherJob {
	var answer []model.FetcherJob

	for _, item := range message.Attch {
		if item.URL == "" {
			continue
		}

		url, err := url.Parse(item.URL)
		if err != nil {
			log.Warn(err)
			continue
		}

		job := model.FetcherJob {
			URL:       url,
			Title:     item.Title,
			OwnerID:   message.UserID,
			ChannelID: message.ChannelID,
			Timestamp: message.DateTS,
		}

		answer = append(answer, job)
	}

	return answer
}

func genericMessage2Jobs(message *model.ArchivedMessage) []model.FetcherJob {
	var answer []model.FetcherJob

    links := xurls.Strict.FindAllString(message.Text, -1);

    for _, link := range links {
        url, err := url.Parse(link)
		if err != nil {
            // strange format in the slack logs, trying to recover
            pieces := strings.Split(link, "|")
            url, err = url.Parse(pieces[0])
            if err != nil {
                log.Warn(err)
    			continue
            }
		}

        job := model.FetcherJob {
            URL: url,
            OwnerID: message.UserID,
            ChannelID: message.ChannelID,
            Timestamp: message.DateTS,
        }

        answer = append(answer, job);
    }

	return answer
}

func fileShares2Jobs(message *model.ArchivedMessage) []model.FetcherJob {
	var answer []model.FetcherJob

	var trustedURL string

	if !message.File.IsExternal {
		trustedURL = message.File.URL
	} else if message.File.LargeCopyURL != "" {
		trustedURL = message.File.LargeCopyURL
	} else if message.File.MediumCopyURL != "" {
		trustedURL = message.File.MediumCopyURL
	} else if message.File.SmallCopyURL != "" {
		trustedURL = message.File.SmallCopyURL
	} else {
		trustedURL = message.File.URL
	}

	url, err := url.Parse(trustedURL)
	if err != nil {
		log.Warn(err)
		return answer
	}

	job := model.FetcherJob {
		URL:       url,
		Title:     message.File.Title,
		MimeType:  message.File.Mime,
		OwnerID:   message.UserID,
		ChannelID: message.ChannelID,
		Timestamp: message.DateTS,
	}

	answer = append(answer, job)

	return answer
}
