package xiami

// AlibabaXiamiApiSearchLetterGetStruct 
type AlibabaXiamiApiSearchLetterGetStruct struct {
    // 艺人结果集
    Artists   []AlibabaXiamiApiSearchLetterGetStruct `json:"artists,omitempty" xml:"artists>alibaba_xiami_api_search_letter_get_struct,omitempty"`
    // 歌曲列表
    Songs   []AlibabaXiamiApiSearchLetterGetStruct `json:"songs,omitempty" xml:"songs>alibaba_xiami_api_search_letter_get_struct,omitempty"`
    // 艺人姓名
    ArtistName   string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
    // 艺人ID
    ArtistId   int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
    // 艺人LOGO
    ArtistLogo   string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
    // 别名
    EnglishName   string `json:"english_name,omitempty" xml:"english_name,omitempty"`
    // 专辑ID
    AlbumId   int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
    // 歌曲ID
    SongId   int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
    // 歌曲名称
    SongName   string `json:"song_name,omitempty" xml:"song_name,omitempty"`
    // 歌手名称
    Singer   string `json:"singer,omitempty" xml:"singer,omitempty"`
    // 专辑名称
    AlbumName   string `json:"album_name,omitempty" xml:"album_name,omitempty"`
    // 专辑LOGO
    AlbumLogo   string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
}
