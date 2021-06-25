package xiami

// StandardArtist 
type StandardArtist struct {

    // 艺人id
    ArtistId   int64 `json:"artist_id,omitempty"`

    // 所属公司
    Company   string `json:"company,omitempty"`

    // 地区
    Area   string `json:"area,omitempty"`

    // M表示男性, F表示女性
    Gender   string `json:"gender,omitempty"`

    // 描述
    Description   string `json:"description,omitempty"`

    // 英文名
    EnglishName   string `json:"english_name,omitempty"`

    // 类别
    Category   int64 `json:"category,omitempty"`

    // 艺人头像
    Logo   string `json:"logo,omitempty"`

    // 专辑数量
    AlbumsCount   int64 `json:"albums_count,omitempty"`

    // 推荐数
    Recommends   int64 `json:"recommends,omitempty"`

    // 试听数
    PlayCount   int64 `json:"play_count,omitempty"`

    // 艺人名
    ArtistName   string `json:"artist_name,omitempty"`

}
