package xiami

// AlbumDetail 
type AlbumDetail struct {

    // 专辑类型
    
    AlbumCategory   string `json:"album_category,omitempty" xml:"album_category,omitempty"`
    

    // 专辑评分
    
    Grade   int64 `json:"grade,omitempty" xml:"grade,omitempty"`
    

    // 评论数
    
    Comments   int64 `json:"comments,omitempty" xml:"comments,omitempty"`
    

    // 艺人名
    
    ArtistName   string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
    

    // 专辑ID
    
    AlbumId   int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
    

    // 艺人ID
    
    ArtistId   int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
    

    // 专辑介绍
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 专辑名
    
    AlbumName   string `json:"album_name,omitempty" xml:"album_name,omitempty"`
    

    // 专辑包含歌曲数目
    
    SongCount   int64 `json:"song_count,omitempty" xml:"song_count,omitempty"`
    

    // 语言类型
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

    // 发布的时间戳
    
    GmtPublish   int64 `json:"gmt_publish,omitempty" xml:"gmt_publish,omitempty"`
    

    // 发行公司
    
    Company   string `json:"company,omitempty" xml:"company,omitempty"`
    

    // 专辑别名
    
    SubTitle   string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
    

    // 上下架信息(3为下架 禁止播放)
    
    IsCheck   int64 `json:"is_check,omitempty" xml:"is_check,omitempty"`
    

    // 类型
    
    Category   int64 `json:"category,omitempty" xml:"category,omitempty"`
    

    // CD碟数
    
    CdCount   int64 `json:"cd_count,omitempty" xml:"cd_count,omitempty"`
    

    // 专辑LOGO
    
    Logo   string `json:"logo,omitempty" xml:"logo,omitempty"`
    

    // 艺人LOGO
    
    ArtistLogo   string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
    

    // is_play
    
    IsPlay   int64 `json:"is_play,omitempty" xml:"is_play,omitempty"`
    

    // 推荐数
    
    Recommends   int64 `json:"recommends,omitempty" xml:"recommends,omitempty"`
    

    // 收藏数
    
    Collects   int64 `json:"collects,omitempty" xml:"collects,omitempty"`
    

    // status
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 歌曲码率是否检查
    
    CheckRate   bool `json:"check_rate,omitempty" xml:"check_rate,omitempty"`
    

    // 是否可以播放
    
    PlayAuthority   int64 `json:"play_authority,omitempty" xml:"play_authority,omitempty"`
    

    // 总歌曲播放数
    
    PlayCount   int64 `json:"play_count,omitempty" xml:"play_count,omitempty"`
    

    // 歌曲列表
    
    Songs   []Song `json:"songs,omitempty" xml:"songs,omitempty"`
    

}
