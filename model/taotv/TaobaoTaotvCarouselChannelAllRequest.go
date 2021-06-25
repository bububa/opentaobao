package taotv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有频道列表 APIRequest
taobao.taotv.carousel.channel.all

获取所有频道列表，按照序号升序
*/
type TaobaoTaotvCarouselChannelAllRequest struct {
    model.Params

    // 系统信息
    systemInfo   string 

}

func NewTaobaoTaotvCarouselChannelAllRequest() *TaobaoTaotvCarouselChannelAllRequest{
    return &TaobaoTaotvCarouselChannelAllRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaotvCarouselChannelAllRequest) GetApiMethodName() string {
    return "taobao.taotv.carousel.channel.all"
}

func (r TaobaoTaotvCarouselChannelAllRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaotvCarouselChannelAllRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r TaobaoTaotvCarouselChannelAllRequest) GetSystemInfo() string {
    return r.systemInfo
}

