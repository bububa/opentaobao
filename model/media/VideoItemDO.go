package media

// VideoItemDo 
type VideoItemDo struct {
    // 视频封面
    CoverUrl   string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
    // 视频时长
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    // 视频上传视频
    UploadTime   string `json:"upload_time,omitempty" xml:"upload_time,omitempty"`
    // 视频标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
}
