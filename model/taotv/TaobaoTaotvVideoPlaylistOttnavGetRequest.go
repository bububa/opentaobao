package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ott播单 API请求
taobao.taotv.video.playlist.ottnav.get

根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
*/
type TaobaoTaotvVideoPlaylistOttnavGetRequest struct {
    model.Params
    // 播单id
    playListId   int64
    // 播单列表
    playListNav   []string
    // 系统信息
    systemInfo   string
}

// 初始化TaobaoTaotvVideoPlaylistOttnavGetRequest对象
func NewTaobaoTaotvVideoPlaylistOttnavGetRequest() *TaobaoTaotvVideoPlaylistOttnavGetRequest{
    return &TaobaoTaotvVideoPlaylistOttnavGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.ottnav.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlayListId Setter
// 播单id
func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetPlayListId(playListId int64) error {
    r.playListId = playListId
    r.Set("play_list_id", playListId)
    return nil
}

// PlayListId Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetPlayListId() int64 {
    return r.playListId
}
// PlayListNav Setter
// 播单列表
func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetPlayListNav(playListNav []string) error {
    r.playListNav = playListNav
    r.Set("play_list_nav", playListNav)
    return nil
}

// PlayListNav Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetPlayListNav() []string {
    return r.playListNav
}
// SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetSystemInfo() string {
    return r.systemInfo
}
