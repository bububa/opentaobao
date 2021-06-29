package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-看台接口pushStand APIRequest
alibaba.damai.mev.open.pushstand

pushStand
*/
type AlibabaDamaiMevOpenPushstandRequest struct {
    model.Params

    // 入参pushStandParam
    pushStandParam   *ThirdStandPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushstandRequest() *AlibabaDamaiMevOpenPushstandRequest{
    return &AlibabaDamaiMevOpenPushstandRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushstandRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushstand"
}

func (r AlibabaDamaiMevOpenPushstandRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushstandRequest) SetPushStandParam(pushStandParam *ThirdStandPushOpenParam) error {
    r.pushStandParam = pushStandParam
    r.Set("push_stand_param", pushStandParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushstandRequest) GetPushStandParam() *ThirdStandPushOpenParam {
    return r.pushStandParam
}

