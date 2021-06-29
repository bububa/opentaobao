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
    _playListId   int64
    // 播单列表
    _playListNav   []string
    // 系统信息
    _systemInfo   string
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
func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetPlayListId(_playListId int64) error {
    r._playListId = _playListId
    r.Set("play_list_id", _playListId)
    return nil
}

// PlayListId Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetPlayListId() int64 {
    return r._playListId
}
// PlayListNav Setter
// 播单列表
func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetPlayListNav(_playListNav []string) error {
    r._playListNav = _playListNav
    r.Set("play_list_nav", _playListNav)
    return nil
}

// PlayListNav Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetPlayListNav() []string {
    return r._playListNav
}
// SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetSystemInfo() string {
    return r._systemInfo
}
