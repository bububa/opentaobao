package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据频道ID获取频道下节目单以及当前播放 API请求
taobao.taotv.carousel.playlist.get

根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
*/
type TaobaoTaotvCarouselPlaylistGetAPIRequest struct {
    model.Params
    // 频道ID
    _channelId   int64
    // 设备信息
    _systemInfo   string
}

// 初始化TaobaoTaotvCarouselPlaylistGetAPIRequest对象
func NewTaobaoTaotvCarouselPlaylistGetRequest() *TaobaoTaotvCarouselPlaylistGetAPIRequest{
    return &TaobaoTaotvCarouselPlaylistGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetApiMethodName() string {
    return "taobao.taotv.carousel.playlist.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelId Setter
// 频道ID
func (r *TaobaoTaotvCarouselPlaylistGetAPIRequest) SetChannelId(_channelId int64) error {
    r._channelId = _channelId
    r.Set("channel_id", _channelId)
    return nil
}

// ChannelId Getter
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetChannelId() int64 {
    return r._channelId
}
// SystemInfo Setter
// 设备信息
func (r *TaobaoTaotvCarouselPlaylistGetAPIRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvCarouselPlaylistGetAPIRequest) GetSystemInfo() string {
    return r._systemInfo
}
