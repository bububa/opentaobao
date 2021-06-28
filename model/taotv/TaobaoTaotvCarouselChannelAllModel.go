package taotv

// TaobaoTaotvCarouselChannelAllModel 
/* model for simplify = false
type TaobaoTaotvCarouselChannelAllModel struct {

    // 频道固定的编号
    
    SerialNumber   int64 `json:"serial_number,omitempty"`
    

    // 频道所有的视频列表
    
    VideoList  struct {
        Videolist  []Videolist `json:"videolist,omitempty"`
    } `json:"video_list,omitempty"`
    

    // 频道当前播放视频
    
    CurrentVideo  *struct {
        CarouselCurrentPlayRbo  *CarouselCurrentPlayRbo `json:"carousel_current_play_rbo,omitempty"`
    } `json:"current_video,omitempty"`
    

    // 频道描述
    
    Description   string `json:"description,omitempty"`
    

    // 牌照方
    
    Bcp   int64 `json:"bcp,omitempty"`
    

    // 频道图标
    
    Pic   string `json:"pic,omitempty"`
    

    // 频道名称
    
    Name   string `json:"name,omitempty"`
    

    // 频道ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 切换时间点
    
    SwitchTime   int64 `json:"switch_time,omitempty"`
    

}
*/

// TaobaoTaotvCarouselChannelAllModel 
type TaobaoTaotvCarouselChannelAllModel struct {

    // 频道固定的编号
    SerialNumber   int64 `json:"serial_number,omitempty"`

    // 频道所有的视频列表
    VideoList   []Videolist `json:"video_list,omitempty"`

    // 频道当前播放视频
    CurrentVideo   *CarouselCurrentPlayRbo `json:"current_video,omitempty"`

    // 频道描述
    Description   string `json:"description,omitempty"`

    // 牌照方
    Bcp   int64 `json:"bcp,omitempty"`

    // 频道图标
    Pic   string `json:"pic,omitempty"`

    // 频道名称
    Name   string `json:"name,omitempty"`

    // 频道ID
    Id   int64 `json:"id,omitempty"`

    // 切换时间点
    SwitchTime   int64 `json:"switch_time,omitempty"`

}
