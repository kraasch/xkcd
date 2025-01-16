package types

const (
    URL_START = "https://xkcd.com/"
    URL_END   = "/info.0.json"
    DATA_DIR  = "./.data"
)

type XKCDComic struct {
	Title string `json:"title"`
	Alt   string `json:"alt"`
	IMG   string `json:"img"`
}

