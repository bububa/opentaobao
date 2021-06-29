package aliexpress

// OffsitePostVideoVO 
type OffsitePostVideoVO struct {

    // 视频封面，与视频尺寸要一致
    
    CoverUrl   string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
    

    // 视频地址
    
    VideoUrl   string `json:"video_url,omitempty" xml:"video_url,omitempty"`
    

    // 视频宽度
    
    Width   int64 `json:"width,omitempty" xml:"width,omitempty"`
    

    // 视频高度
    
    Height   int64 `json:"height,omitempty" xml:"height,omitempty"`
    

}
