package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据频道ID获取频道下节目单以及当前播放 APIRequest
taobao.taotv.carousel.playlist.get

根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
*/
type TaobaoTaotvCarouselPlaylistGetRequest struct {
    model.Params

    // 频道ID
    channelId   int64 

    // 设备信息
    systemInfo   string 

}

func NewTaobaoTaotvCarouselPlaylistGetRequest() *TaobaoTaotvCarouselPlaylistGetRequest{
    return &TaobaoTaotvCarouselPlaylistGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaotvCarouselPlaylistGetRequest) GetApiMethodName() string {
    return "taobao.taotv.carousel.playlist.get"
}

func (r TaobaoTaotvCarouselPlaylistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaotvCarouselPlaylistGetRequest) SetChannelId(channelId int64) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

func (r TaobaoTaotvCarouselPlaylistGetRequest) GetChannelId() int64 {
    return r.channelId
}

func (r *TaobaoTaotvCarouselPlaylistGetRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r TaobaoTaotvCarouselPlaylistGetRequest) GetSystemInfo() string {
    return r.systemInfo
}

