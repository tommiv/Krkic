package di

import (
    "os"
    "path"
    "gopkg.in/spf13/viper.v0"
)

func EnsureFolders() {
    root := viper.GetString("app_data_folder")

    os.MkdirAll(path.Join(root, "imgcache"), os.ModePerm)
}
