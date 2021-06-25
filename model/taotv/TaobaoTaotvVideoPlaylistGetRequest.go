package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据频道ID获取频道下节目单以及当前播放 APIRequest
taobao.taotv.video.playlist.get

根据频道ID获取频道下节目单以及当前播放
*/
type TaobaoTaotvVideoPlaylistGetRequest struct {
    model.Params

    // 播单id
    playListId   int64 

    // 系统信息
    systemInfo   string 

}

func NewTaobaoTaotvVideoPlaylistGetRequest() *TaobaoTaotvVideoPlaylistGetRequest{
    return &TaobaoTaotvVideoPlaylistGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaotvVideoPlaylistGetRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.get"
}

func (r TaobaoTaotvVideoPlaylistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaotvVideoPlaylistGetRequest) SetPlayListId(playListId int64) error {
    r.playListId = playListId
    r.Set("play_list_id", playListId)
    return nil
}

func (r TaobaoTaotvVideoPlaylistGetRequest) GetPlayListId() int64 {
    return r.playListId
}

func (r *TaobaoTaotvVideoPlaylistGetRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r TaobaoTaotvVideoPlaylistGetRequest) GetSystemInfo() string {
    return r.systemInfo
}

