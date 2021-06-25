package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取播单列表 APIRequest
taobao.taotv.video.playlist.all

根据牌照和视频源等获取播单列表
*/
type TaobaoTaotvVideoPlaylistAllRequest struct {
    model.Params

    // 系统信息
    systemInfo   string 

}

func NewTaobaoTaotvVideoPlaylistAllRequest() *TaobaoTaotvVideoPlaylistAllRequest{
    return &TaobaoTaotvVideoPlaylistAllRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaotvVideoPlaylistAllRequest) GetApiMethodName() string {
    return "taobao.taotv.video.playlist.all"
}

func (r TaobaoTaotvVideoPlaylistAllRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaotvVideoPlaylistAllRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r TaobaoTaotvVideoPlaylistAllRequest) GetSystemInfo() string {
    return r.systemInfo
}

