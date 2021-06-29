package alihouse

// VideoDataDTO 
type VideoDataDTO struct {

    // 视频下载链接
    
    VideoUrl   string `json:"video_url,omitempty" xml:"video_url,omitempty"`
    

    // 视频封面
    
    CoverImage   string `json:"cover_image,omitempty" xml:"cover_image,omitempty"`
    

    // 外部视频id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

}
