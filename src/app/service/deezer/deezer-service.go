package deezer

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"search/src/config"
	"time"
)

type Client struct {
	q string
	c config.Config
}

type Response struct {
	Data          []Record `json:"data"`
	Total         int      `json:"total"`
	NextPageToken string   `json:"next"`
}

type Record struct {
	ID                    int    `json:"id"`
	Readable              bool   `json:"readable"`
	Title                 string `json:"title"`
	TitleShort            string `json:"title_short"`
	Link                  string `json:"link"`
	Duration              int    `json:"duration"`
	Rank                  int    `json:"rank"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
	ExplicitContentCover  int    `json:"explicit_content_cover"`
	Preview               string `json:"preview"`
	ArtistInfo            Artist `json:"artist"`
	AlbumInfo             Album  `json:"album"`
	Type                  string `json:"type"`
}

type Artist struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXL     string `json:"picture_xl"`
	TrackList     string `json:"tracklist"`
	Type          string `json:"type"`
}

type Album struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	CoverSmall  string `json:"cover_small"`
	CoverMedium string `json:"cover_medium"`
	CoverBig    string `json:"cover_big"`
	CoverXL     string `json:"cover_xl"`
	TrackList   string `json:"tracklist"`
	Type        string `json:"type"`
}

func NewClient(conf config.Config, query string) *Client {
	return &Client{
		q: query,
		c: conf,
	}
}

func (c *Client) BuildRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", c.c.DeezerConfig.ApiFqdn+url.QueryEscape(c.q), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(c.c.DeezerConfig.ApiHostHeader, c.c.DeezerConfig.ApiHost)
	req.Header.Add(c.c.DeezerConfig.ApiKeyHeader, c.c.DeezerConfig.ApiKey)
	return req, nil
}

func (c *Client) Execute(r *http.Request) (*Response, error) {
	client := buildClient()
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else if string(bodyBytes) == "" {
			return nil, fmt.Errorf("error while excuting search request for query [%s], status code [%d]", c.q, resp.StatusCode)
		}
		return nil, fmt.Errorf(string(bodyBytes))
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	}
}

func buildClient() *http.Client {
	to := 10 * time.Second
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			KeepAlive: to,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{
		Transport: tr,
		Timeout:   to,
	}
}
