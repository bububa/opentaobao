package xiami

// StandardSong 
/* model for simplify = false
type StandardSong struct {

    // 试听文件地址
    
    ListenFile   string `json:"listen_file,omitempty"`
    

    // 歌曲id
    
    SongId   int64 `json:"song_id,omitempty"`
    

    // 歌曲名
    
    SongName   string `json:"song_name,omitempty"`
    

    // 专辑名
    
    AlbumName   string `json:"album_name,omitempty"`
    

    // 专辑id
    
    AlbumId   int64 `json:"album_id,omitempty"`
    

    // 艺人名字
    
    ArtistName   string `json:"artist_name,omitempty"`
    

    // 艺人id
    
    ArtistId   int64 `json:"artist_id,omitempty"`
    

    // 推荐数
    
    Recommends   int64 `json:"recommends,omitempty"`
    

    // 歌词地址
    
    Lyric   string `json:"lyric,omitempty"`
    

    // 歌曲长度(秒)
    
    Length   int64 `json:"length,omitempty"`
    

    // 专辑封面
    
    Logo   string `json:"logo,omitempty"`
    

    // 演唱者
    
    Singers   string `json:"singers,omitempty"`
    

    // 试听数
    
    PlayCounts   int64 `json:"play_counts,omitempty"`
    

}
*/

// StandardSong 
type StandardSong struct {

    // 试听文件地址
    ListenFile   string `json:"listen_file,omitempty"`

    // 歌曲id
    SongId   int64 `json:"song_id,omitempty"`

    // 歌曲名
    SongName   string `json:"song_name,omitempty"`

    // 专辑名
    AlbumName   string `json:"album_name,omitempty"`

    // 专辑id
    AlbumId   int64 `json:"album_id,omitempty"`

    // 艺人名字
    ArtistName   string `json:"artist_name,omitempty"`

    // 艺人id
    ArtistId   int64 `json:"artist_id,omitempty"`

    // 推荐数
    Recommends   int64 `json:"recommends,omitempty"`

    // 歌词地址
    Lyric   string `json:"lyric,omitempty"`

    // 歌曲长度(秒)
    Length   int64 `json:"length,omitempty"`

    // 专辑封面
    Logo   string `json:"logo,omitempty"`

    // 演唱者
    Singers   string `json:"singers,omitempty"`

    // 试听数
    PlayCounts   int64 `json:"play_counts,omitempty"`

}
