package xiamiopen

// SongDetailDo 
type SongDetailDo struct {

    // mvId
    
    MvId   string `json:"mv_id,omitempty" xml:"mv_id,omitempty"`
    

    // 专辑信息
    
    Album   *AlbumDo `json:"album,omitempty" xml:"album,omitempty"`
    

    // 专辑id
    
    AlbumId   int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
    

    // 专辑logo
    
    AlbumLogo   string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
    

    // 编曲
    
    Arrangement   string `json:"arrangement,omitempty" xml:"arrangement,omitempty"`
    

    // artistDO
    
    Artist   *ArtistDo `json:"artist,omitempty" xml:"artist,omitempty"`
    

    // 艺人id
    
    ArtistId   int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
    

    // cd序号
    
    CdSerial   int64 `json:"cd_serial,omitempty" xml:"cd_serial,omitempty"`
    

    // 作词
    
    Composer   string `json:"composer,omitempty" xml:"composer,omitempty"`
    

    // 演唱者
    
    Singers   []ArtistDo `json:"singers,omitempty" xml:"singers,omitempty"`
    

    // 歌曲id
    
    SongId   int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
    

    // 歌曲名
    
    SongName   string `json:"song_name,omitempty" xml:"song_name,omitempty"`
    

    // 作曲
    
    Songwriters   string `json:"songwriters,omitempty" xml:"songwriters,omitempty"`
    

    // 序号
    
    Track   int64 `json:"track,omitempty" xml:"track,omitempty"`
    

    // 歌词文件地址
    
    LyricFile   string `json:"lyric_file,omitempty" xml:"lyric_file,omitempty"`
    

    // 歌词类型
    
    LyricType   int64 `json:"lyric_type,omitempty" xml:"lyric_type,omitempty"`
    

    // 版权信息
    
    PurviewInfo   string `json:"purview_info,omitempty" xml:"purview_info,omitempty"`
    

    // 歌曲时长
    
    Length   int64 `json:"length,omitempty" xml:"length,omitempty"`
    

}
