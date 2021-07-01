package xiamicontent

// SearchTermsDto 结构体
type SearchTermsDto struct {
	// 搜索的value。songName:歌曲名称；singerName:演唱者名称；lyric:歌词文本（只匹配前50字符）； copyrightStatus：版权状态 0下架/1上架；publishStatus：发布状态 0未发布/1发布; keyword:关键字搜索（与songName/singerName互斥）; lyric:歌词搜索，限制字符数50
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 搜索key：songName/singerName/lyric/copyrightStatus/publishStatus/keyword/lyric
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
