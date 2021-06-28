package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ott播单 APIRequest
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

func NewTaobaoTaotvVideoPlaylistOttnavGetRequest() *TaobaoTaotvVideoPlaylistOttnavGetRequest{
    return &TaobaoTaotvVideoPlaylistOttnavGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.ottnav.get"
}

func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetPlayListId(playListId int64) error {
    r.playListId = playListId
    r.Set("play_list_id", playListId)
    return nil
}

func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetPlayListId() int64 {
    return r.playListId
}

func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetPlayListNav(playListNav []string) error {
    r.playListNav = playListNav
    r.Set("play_list_nav", playListNav)
    return nil
}

func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetPlayListNav() []string {
    return r.playListNav
}

func (r *TaobaoTaotvVideoPlaylistOttnavGetRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r TaobaoTaotvVideoPlaylistOttnavGetRequest) GetSystemInfo() string {
    return r.systemInfo
}

