package parser

import "search/src/app/service/deezer"

type ParsedResponse struct {
	Data []ParsedRecord `json:"data"`
}

type ParsedRecord struct {
	Title      string       `json:"title"`
	Link       string       `json:"link"`
	Preview    string       `json:"preview"`
	ArtistInfo ParsedArtist `json:"artist"`
	AlbumInfo  ParsedAlbum  `json:"album"`
}

type ParsedArtist struct {
	Name    string `json:"name"`
	Link    string `json:"link"`
	Picture string `json:"picture"`
}

type ParsedAlbum struct {
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	TrackList string `json:"tracklist"`
}

func ParseResponse(resp *deezer.Response) *ParsedResponse {
	var data []ParsedRecord
	for _, r := range resp.Data {
		pa := ParsedArtist{
			Name:    r.ArtistInfo.Name,
			Link:    r.ArtistInfo.Link,
			Picture: r.ArtistInfo.Picture,
		}

		pal := ParsedAlbum{
			Title:     r.AlbumInfo.Title,
			Cover:     r.AlbumInfo.Cover,
			TrackList: r.AlbumInfo.TrackList,
		}

		pr := ParsedRecord{
			Title:      r.Title,
			Link:       r.Link,
			Preview:    r.Preview,
			ArtistInfo: pa,
			AlbumInfo:  pal,
		}

		data = append(data, pr)
	}
	return &ParsedResponse{Data: data}
}