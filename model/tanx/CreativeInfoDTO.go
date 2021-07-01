package tanx

// CreativeInfoDto 
type CreativeInfoDto struct {
    // dsp系统中的创意id
    CreativeId   string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
    // 创意尺寸，长高中间用小写英文x
    CreativeSize   string `json:"creative_size,omitempty" xml:"creative_size,omitempty"`
    // 创意的类目，多个值用逗号&ldquo;，&rdquo;分隔
    CreativeCategoryId   string `json:"creative_category_id,omitempty" xml:"creative_category_id,omitempty"`
    // 创意类型 1:文字 2:图片 3:Flash 4:视频贴片
    CreativeFormat   string `json:"creative_format,omitempty" xml:"creative_format,omitempty"`
    // 托管创意的名称
    CreativeName   string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
    // 创意支持的apiFramework协议,1:VPAID1.0;2:VPAID2.0;3:MARID-1;4:ORMMA;5:1MRAID-2,只能单选，不能多选
    ApiFramework   string `json:"api_framework,omitempty" xml:"api_framework,omitempty"`
    // 文件格式, 视频类型：flv、avi、mp4, 图片类型：jpg、png、gif, flash类型：swf等
    FileType   string `json:"file_type,omitempty" xml:"file_type,omitempty"`
    // 创意存储地址
    CreativeUrl   string `json:"creative_url,omitempty" xml:"creative_url,omitempty"`
    // 广告跳转的最终目标页面地址
    ClickUrl   string `json:"click_url,omitempty" xml:"click_url,omitempty"`
    // 创意时长，单位是毫秒
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    // 广告主Id，多值使用&ldquo;,&rdquo;分隔
    AdvertiserIds   string `json:"advertiser_ids,omitempty" xml:"advertiser_ids,omitempty"`
    // (选填)中间跳转地址（比如DSP的点击服务器地址），但跳转后的最终地址必须和click_url一致。
    ClickThroughUrl   string `json:"click_through_url,omitempty" xml:"click_through_url,omitempty"`
    // (选填)点击监控地址，最多3个，多个值用逗号&ldquo;，&rdquo;分隔
    ClickTrackUrl   string `json:"click_track_url,omitempty" xml:"click_track_url,omitempty"`
}
