package taotv

// CarouselCurrentPlayRbo 
/* model for simplify = false
type CarouselCurrentPlayRbo struct {

    // 当前轮播视频信息
    
    Video  *struct {
        CarouselPlaylistVideoRbo  *CarouselPlaylistVideoRbo `json:"carousel_playlist_video_rbo,omitempty"`
    } `json:"video,omitempty"`
    

    // 当前视频正在播放的时间点（单位秒）
    
    Point   int64 `json:"point,omitempty"`
    

}
*/

// CarouselCurrentPlayRbo 
type CarouselCurrentPlayRbo struct {

    // 当前轮播视频信息
    Video   *CarouselPlaylistVideoRbo `json:"video,omitempty"`

    // 当前视频正在播放的时间点（单位秒）
    Point   int64 `json:"point,omitempty"`

}
