package taotv

// CarouselCurrentPlayRbo 
type CarouselCurrentPlayRbo struct {

    // 当前轮播视频信息
    Video   *CarouselPlaylistVideoRbo `json:"video,omitempty"`

    // 当前视频正在播放的时间点（单位秒）
    Point   int64 `json:"point,omitempty"`

}
