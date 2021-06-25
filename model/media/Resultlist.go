package media

// Resultlist 
type Resultlist struct {

    // 互动id
    Id   int64 `json:"id,omitempty"`

    // 视频链接，需要unicode转换
    VideoUrl   string `json:"video_url,omitempty"`

    // 封面图url
    CoverUrl   string `json:"cover_url,omitempty"`

    // 视频时长 单位秒
    Duration   int64 `json:"duration,omitempty"`

    // 视频标题
    Title   string `json:"title,omitempty"`

    // 视频状态，1有效；0删除
    VideoStatus   int64 `json:"video_status,omitempty"`

}
