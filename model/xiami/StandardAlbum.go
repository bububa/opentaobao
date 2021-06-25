package xiami

// StandardAlbum 
type StandardAlbum struct {

    // 专辑评分
    Grade   int64 `json:"grade,omitempty"`

    // 艺人名
    ArtistName   string `json:"artist_name,omitempty"`

    // 专辑ID
    AlbumId   int64 `json:"album_id,omitempty"`

    // 艺人ID
    ArtistId   int64 `json:"artist_id,omitempty"`

    // 专辑介绍
    Description   string `json:"description,omitempty"`

    // 专辑名
    AlbumName   string `json:"album_name,omitempty"`

    // 专辑包含歌曲数目
    SongCount   int64 `json:"song_count,omitempty"`

    // 语言类型
    Language   string `json:"language,omitempty"`

    // 发布的时间戳
    GmtPublish   int64 `json:"gmt_publish,omitempty"`

    // 发行公司
    Company   string `json:"company,omitempty"`

    // 专辑别名
    SubTitle   string `json:"sub_title,omitempty"`

    // 上下架信息(3为下架 禁止播放)
    IsCheck   int64 `json:"is_check,omitempty"`

    // 类型
    Category   int64 `json:"category,omitempty"`

    // CD碟数
    CdCount   int64 `json:"cd_count,omitempty"`

    // 专辑LOGO
    Logo   string `json:"logo,omitempty"`

    // 推荐数
    Recommends   int64 `json:"recommends,omitempty"`

    // 收藏数
    Collects   int64 `json:"collects,omitempty"`

}
