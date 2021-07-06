package xiami

// AlibabaXiamiApiArtistMusiclistGetStruct 结构体
type AlibabaXiamiApiArtistMusiclistGetStruct struct {
	// 艺人信息
	Artists []AlibabaXiamiApiArtistMusiclistGetStruct `json:"artists,omitempty" xml:"artists>alibaba_xiami_api_artist_musiclist_get_struct,omitempty"`
	// 艺人名字
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 艺人头像
	ArtistLogo string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
	// 拼音
	Pinyin string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
	// 粉丝数
	CountLikes int64 `json:"count_likes,omitempty" xml:"count_likes,omitempty"`
	// 艺人id, BIGINT类型
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
	// 艺人是否有演出的标记
	IsShow bool `json:"is_show,omitempty" xml:"is_show,omitempty"`
	// 是否是音乐人
	IsMusician bool `json:"is_musician,omitempty" xml:"is_musician,omitempty"`
}
