package xiamicontent

// SongInfoDto 
type SongInfoDto struct {
    // 演唱者列表
    Singers   []ArtistDto `json:"singers,omitempty" xml:"singers>artist_dto,omitempty"`
    // 歌曲名
    SongName   string `json:"song_name,omitempty" xml:"song_name,omitempty"`
    // 作曲人列表
    Composers   []ArtistDto `json:"composers,omitempty" xml:"composers>artist_dto,omitempty"`
    // 专辑艺人
    Artists   []ArtistDto `json:"artists,omitempty" xml:"artists>artist_dto,omitempty"`
    // 专辑信息
    Album   *AlbumDto `json:"album,omitempty" xml:"album,omitempty"`
    // 作词人列表
    Songwriters   []ArtistDto `json:"songwriters,omitempty" xml:"songwriters>artist_dto,omitempty"`
    // 制作人列表
    Producers   []ArtistDto `json:"producers,omitempty" xml:"producers>artist_dto,omitempty"`
    // 标签列表
    Tags   []TagLink `json:"tags,omitempty" xml:"tags>tag_link,omitempty"`
    // 歌曲时长（单位毫秒），没有对应的音频文件时为0
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    // 版权状态
    CopyrightStatus   int64 `json:"copyright_status,omitempty" xml:"copyright_status,omitempty"`
    // 编曲人列表
    Arrangements   []ArtistDto `json:"arrangements,omitempty" xml:"arrangements>artist_dto,omitempty"`
    // 歌词列表
    Lyrics   []LyricDto `json:"lyrics,omitempty" xml:"lyrics>lyric_dto,omitempty"`
    // 歌曲副标题
    SongSubName   string `json:"song_sub_name,omitempty" xml:"song_sub_name,omitempty"`
    // 歌曲id
    SongId   int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
    // 歌曲发布状态, 1-已发布, 0-未发布
    PublishStatus   int64 `json:"publish_status,omitempty" xml:"publish_status,omitempty"`
    // 波形图文件
    WaveformUrl   string `json:"waveform_url,omitempty" xml:"waveform_url,omitempty"`
    // 副歌结束时间戳
    HotPartEndTime   int64 `json:"hot_part_end_time,omitempty" xml:"hot_part_end_time,omitempty"`
    // 副歌开始时间戳
    HotPartStartTime   int64 `json:"hot_part_start_time,omitempty" xml:"hot_part_start_time,omitempty"`
}
