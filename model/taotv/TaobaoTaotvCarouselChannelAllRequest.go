package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有频道列表 API请求
taobao.taotv.carousel.channel.all

获取所有频道列表，按照序号升序
*/
type TaobaoTaotvCarouselChannelAllRequest struct {
    model.Params
    // 系统信息
    _systemInfo   string
}

// 初始化TaobaoTaotvCarouselChannelAllRequest对象
func NewTaobaoTaotvCarouselChannelAllRequest() *TaobaoTaotvCarouselChannelAllRequest{
    return &TaobaoTaotvCarouselChannelAllRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselChannelAllRequest) GetApiMethodName() string {
    return "taobao.taotv.carousel.channel.all"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselChannelAllRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvCarouselChannelAllRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r TaobaoTaotvCarouselChannelAllRequest) GetSystemInfo() string {
    return r._systemInfo
}
