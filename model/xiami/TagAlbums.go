package xiami

// TagAlbums 
type TagAlbums struct {
    // 专辑id
    AlbumId   int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
    // 艺人名
    ArtistName   string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
    // 状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 分数
    Grade   int64 `json:"grade,omitempty" xml:"grade,omitempty"`
    // 艺人id
    ArtistId   int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
    // 专辑名
    AlbumName   string `json:"album_name,omitempty" xml:"album_name,omitempty"`
    // 是否可试听，0：不可，1：可试听
    IsPlay   int64 `json:"is_play,omitempty" xml:"is_play,omitempty"`
    // 专辑封面
    AlbumLogo   string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
}
