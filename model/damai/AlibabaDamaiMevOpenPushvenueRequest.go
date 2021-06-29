package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场馆接口pushVenue APIRequest
alibaba.damai.mev.open.pushvenue

开放接口推送场馆
*/
type AlibabaDamaiMevOpenPushvenueRequest struct {
    model.Params

    // 入参pushVenueParam
    pushVenueParam   *ThirdVenuePushOpenParam 

}

func NewAlibabaDamaiMevOpenPushvenueRequest() *AlibabaDamaiMevOpenPushvenueRequest{
    return &AlibabaDamaiMevOpenPushvenueRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushvenueRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushvenue"
}

func (r AlibabaDamaiMevOpenPushvenueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushvenueRequest) SetPushVenueParam(pushVenueParam *ThirdVenuePushOpenParam) error {
    r.pushVenueParam = pushVenueParam
    r.Set("push_venue_param", pushVenueParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushvenueRequest) GetPushVenueParam() *ThirdVenuePushOpenParam {
    return r.pushVenueParam
}

