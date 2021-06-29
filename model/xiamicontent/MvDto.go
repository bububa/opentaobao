package xiamicontent

// MvDto 
type MvDto struct {
    // 副标题
    SubTitle   string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
    // 标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // mvid
    MvId   string `json:"mv_id,omitempty" xml:"mv_id,omitempty"`
    // 播放地址
    PlayUrl   string `json:"play_url,omitempty" xml:"play_url,omitempty"`
    // 时长（秒）
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    // mv封面地址
    MvCover   string `json:"mv_cover,omitempty" xml:"mv_cover,omitempty"`
    // mv介绍
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
}
