package main

import (
	// "io"
	"krkic/di"
    // "krkic/model"
    // "krkic/fetch"
    "krkic/training"

	"os"
	// "path"

	// "github.com/nlopes/slack"
	// "gopkg.in/spf13/viper.v0"
    log "gopkg.in/Sirupsen/logrus.v0"
)

// TODO: deal with pointer/value type returns everywhere
func main() {
	appDataFolder := os.Getenv("KRKIC_DATA_FOLDER")
	if appDataFolder == "" {
		appDataFolder = "/usr/local/etc/krkic/"
	}

    di.SetupViper(appDataFolder)
    di.SetupLogrus()
    di.SetupRedis()

    messages := training.ReadMessages(appDataFolder)
    log.Info(len(messages))

    jobs := training.ParseJobs(messages);
    log.Info(len(jobs))

    bojans := training.FetchBojans(jobs[0:50])

    for _, bojan := range bojans {
        log.Info(bojan)
    }

    // url := "http://i3.kym-cdn.com/photos/images/facebook/000/862/065/0e9.jpg"
    // data, err := fetch.GetHashByUrl(url)
    // log.Infof("%s, %s", data, err)

	// api := slack.New(viper.GetString("slack_api_key"))
    //
	// channels, err := api.GetChannels(true)
    //
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// 	return
	// }
    //
	// for _, channel := range channels {
	// 	fmt.Printf("ID: %s, Name: %s\n", channel.ID, channel.Name)
	// }
    //
	// rtm := api.NewRTM()
	// go rtm.ManageConnection()
    //
	// for msg := range rtm.IncomingEvents {
	// 	switch ev := msg.Data.(type) {
	// 	case *slack.ConnectedEvent:
	// 		// fmt.Println("Infos:", ev.Info)
	// 		// fmt.Println("Connection counter:", ev.ConnectionCount)
	// 		// Replace #general with your Channel ID
	// 		// rtm.SendMessage(rtm.NewOutgoingMessage("```К И Б Е Р Р Е А Л Т А Й М```", "C3VE6S88N"))
    //
	// 	case *slack.MessageEvent:
	// 		fmt.Printf("Message: %v\n", ev)
    //
	// 	case *slack.LatencyReport:
	// 		fmt.Printf("Current latency: %v\n", ev.Value)
    //
	// 	case *slack.RTMError:
	// 		fmt.Printf("Error: %s\n", ev.Error())
    //
	// 	case *slack.InvalidAuthEvent:
	// 		fmt.Printf("Invalid credentials")
	// 		return
    //
	// 	default:
	// 		// Ignore other events..
	// 		// fmt.Printf("Unexpected: %v\n", msg.Data)
	// 	}
	// }
}
