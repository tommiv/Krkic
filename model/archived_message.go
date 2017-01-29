package model

type ArchivedFile struct {
    Name          string  `json:"name"`
    Title         string  `json:"title"`
    Mime          string  `json:"mimetype"`
    Extension     string  `json:"filetype"`
    IsExternal    bool    `json:"is_external"`
    URL           string  `json:"url_private"`
    LargeCopyURL  string  `json:"thumb_720"`
    MediumCopyURL string  `json:"thumb_480"`
    SmallCopyURL  string  `json:"thumb_360"`
}

type ArchivedAttch struct {
    Title string `json:"title"`
    URL   string `json:"from_url"`
}

type ArchivedMessage struct {
    UserID   string          `json:"user"`
    Type     string          `json:"type"`
    Subtype  string          `json:"subtype"`
    Text     string          `json:"text"`
    DateTS   string          `json:"ts"`
    File     ArchivedFile    `json:"file"`
    Attch    []ArchivedAttch `json:"attachments"`
}
