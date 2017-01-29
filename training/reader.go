package training

import (
    "path"
    "io/ioutil"
    "encoding/json"
    "krkic/model"
    log "gopkg.in/Sirupsen/logrus.v0"
)

func ReadMessages(appDataFolder string) []model.ArchivedMessage {
    dumpRoot := path.Join(appDataFolder, "dump")

    channelsPath := path.Join(dumpRoot, "channels.json")
    channelsBuf, err := ioutil.ReadFile(channelsPath)
    if err != nil {
        panic(err)
    }

    var channels []map[string]interface{}
    json.Unmarshal(channelsBuf, &channels)

    var messages []model.ArchivedMessage

    for _, channel := range channels {
        chanName := channel["name"].(string)
        chanPath := path.Join(dumpRoot, chanName)

        files, _ := ioutil.ReadDir(chanPath)
        for _, file := range files {
            filePath := path.Join(chanPath, file.Name())
            fileBuf, err := ioutil.ReadFile(filePath)
            if err != nil {
                log.Warningf("Chunk [%s] is unreadable", filePath)
                continue
            }

            var chunk []model.ArchivedMessage
            json.Unmarshal(fileBuf, &chunk)

            messages = append(messages, chunk...)
        }
    }

    return messages
}
