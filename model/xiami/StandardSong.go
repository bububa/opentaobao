package xiami

// StandardSong 
type StandardSong struct {

    // 试听文件地址
    
    ListenFile   string `json:"listen_file,omitempty" xml:"listen_file,omitempty"`
    

    // 歌曲id
    
    SongId   int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
    

    // 歌曲名
    
    SongName   string `json:"song_name,omitempty" xml:"song_name,omitempty"`
    

    // 专辑名
    
    AlbumName   string `json:"album_name,omitempty" xml:"album_name,omitempty"`
    

    // 专辑id
    
    AlbumId   int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
    

    // 艺人名字
    
    ArtistName   string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
    

    // 艺人id
    
    ArtistId   int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
    

    // 推荐数
    
    Recommends   int64 `json:"recommends,omitempty" xml:"recommends,omitempty"`
    

    // 歌词地址
    
    Lyric   string `json:"lyric,omitempty" xml:"lyric,omitempty"`
    

    // 歌曲长度(秒)
    
    Length   int64 `json:"length,omitempty" xml:"length,omitempty"`
    

    // 专辑封面
    
    Logo   string `json:"logo,omitempty" xml:"logo,omitempty"`
    

    // 演唱者
    
    Singers   string `json:"singers,omitempty" xml:"singers,omitempty"`
    

    // 试听数
    
    PlayCounts   int64 `json:"play_counts,omitempty" xml:"play_counts,omitempty"`
    

}
