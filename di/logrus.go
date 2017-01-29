package di

import (
    "io"
    "os"
    "path"
    "gopkg.in/spf13/viper.v0"
    log "gopkg.in/Sirupsen/logrus.v0"
)

func SetupLogrus(appDataFolder string) {
    logFile, err := os.OpenFile(
        path.Join(appDataFolder, "krkic.log"),
        os.O_WRONLY | os.O_CREATE | os.O_APPEND,
        0755,
    )

    if err != nil {
        panic(err)
    }

    level, err := log.ParseLevel(viper.GetString("log_level"))

    if err != nil {
        panic(err)
    }

    logOutput := io.MultiWriter(os.Stdout, logFile)
    log.SetOutput(logOutput)
    log.SetLevel(level)

    formatter := &log.TextFormatter{
        FullTimestamp: true,
    }
    log.SetFormatter(formatter)
}
