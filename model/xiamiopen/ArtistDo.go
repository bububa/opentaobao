package xiamiopen

// ArtistDo 结构体
type ArtistDo struct {
	// 艺人别名
	Alias string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 艺人logo
	ArtistLogo string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
	// 艺人名
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 艺人id
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
	// 是否音乐人
	Musician bool `json:"musician,omitempty" xml:"musician,omitempty"`
}
