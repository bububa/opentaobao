package media

// VideoItemExtDO 
/* model for simplify = false
type VideoItemExtDO struct {

    // 视频播放地址
    
    PlayUrl   string `json:"play_url,omitempty"`
    

    // 视频封面-主图
    
    MainPicUrl   string `json:"main_pic_url,omitempty"`
    

    // 视频状态
    
    State   int64 `json:"state,omitempty"`
    

    // 是否能在移动端播放
    
    CanPlayInPhone   bool `json:"can_play_in_phone,omitempty"`
    

    // 视频基本信息
    
    VideoInfo  *struct {
        VideoItemDO  *VideoItemDO `json:"video_item_do,omitempty"`
    } `json:"video_info,omitempty"`
    

}
*/

// VideoItemExtDO 
type VideoItemExtDO struct {

    // 视频播放地址
    PlayUrl   string `json:"play_url,omitempty"`

    // 视频封面-主图
    MainPicUrl   string `json:"main_pic_url,omitempty"`

    // 视频状态
    State   int64 `json:"state,omitempty"`

    // 是否能在移动端播放
    CanPlayInPhone   bool `json:"can_play_in_phone,omitempty"`

    // 视频基本信息
    VideoInfo   *VideoItemDO `json:"video_info,omitempty"`

}
