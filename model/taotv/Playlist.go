package taotv

// Playlist 
type Playlist struct {
    // 播单ID
    PlayListId   int64 `json:"play_list_id,omitempty" xml:"play_list_id,omitempty"`
    // 播单名称
    PlayListName   string `json:"play_list_name,omitempty" xml:"play_list_name,omitempty"`
    // 背景图
    BgPic   string `json:"bg_pic,omitempty" xml:"bg_pic,omitempty"`
    // icon图标
    IconPic   string `json:"icon_pic,omitempty" xml:"icon_pic,omitempty"`
    // 是否开启视频推荐功能：1开启0关闭（null或无字段均关闭)
    HasRecommend   int64 `json:"has_recommend,omitempty" xml:"has_recommend,omitempty"`
}
