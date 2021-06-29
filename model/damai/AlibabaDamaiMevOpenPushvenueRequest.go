package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场馆接口pushVenue API请求
alibaba.damai.mev.open.pushvenue

开放接口推送场馆
*/
type AlibabaDamaiMevOpenPushvenueRequest struct {
    model.Params
    // 入参pushVenueParam
    _pushVenueParam   *ThirdVenuePushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushvenueRequest对象
func NewAlibabaDamaiMevOpenPushvenueRequest() *AlibabaDamaiMevOpenPushvenueRequest{
    return &AlibabaDamaiMevOpenPushvenueRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushvenueRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushvenue"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushvenueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushVenueParam Setter
// 入参pushVenueParam
func (r *AlibabaDamaiMevOpenPushvenueRequest) SetPushVenueParam(_pushVenueParam *ThirdVenuePushOpenParam) error {
    r._pushVenueParam = _pushVenueParam
    r.Set("push_venue_param", _pushVenueParam)
    return nil
}

// PushVenueParam Getter
func (r AlibabaDamaiMevOpenPushvenueRequest) GetPushVenueParam() *ThirdVenuePushOpenParam {
    return r._pushVenueParam
}
