package iot

// VideoUrlParam 
type VideoUrlParam struct {
    // 视频封面图
    Cover   *ImageUrlParam `json:"cover,omitempty" xml:"cover,omitempty"`
    // 默认播放链接
    DefaultUrl   string `json:"default_url,omitempty" xml:"default_url,omitempty"`
    // 视频高度
    Height   int64 `json:"height,omitempty" xml:"height,omitempty"`
    // 高清
    High   string `json:"high,omitempty" xml:"high,omitempty"`
    // 标清
    Standard   string `json:"standard,omitempty" xml:"standard,omitempty"`
    // 超清
    Ultra   string `json:"ultra,omitempty" xml:"ultra,omitempty"`
    // 视频宽度
    Width   int64 `json:"width,omitempty" xml:"width,omitempty"`
}
