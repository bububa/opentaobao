package iot

// VideoUrlDto 
type VideoUrlDto struct {

    // 视频封面图
    Cover   *ImageUrlDto `json:"cover,omitempty"`

    // 默认播放链接
    DefaultUrl   string `json:"default_url,omitempty"`

    // 视频高度
    Height   int64 `json:"height,omitempty"`

    // 高清
    High   string `json:"high,omitempty"`

    // 标清
    Standard   string `json:"standard,omitempty"`

    // 超清
    Ultra   string `json:"ultra,omitempty"`

    // 视频宽度
    Width   int64 `json:"width,omitempty"`

}
