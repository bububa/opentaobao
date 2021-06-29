package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】开放平台对外页面跳转 APIRequest
alitrip.btrip.openplatform.address.get

获取类目预定页跳转地址
*/
type AlitripBtripOpenplatformAddressGetRequest struct {
    model.Params

    // 入参
    rq   *OpenApiJumpInfoRq 

}

func NewAlitripBtripOpenplatformAddressGetRequest() *AlitripBtripOpenplatformAddressGetRequest{
    return &AlitripBtripOpenplatformAddressGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripOpenplatformAddressGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.openplatform.address.get"
}

func (r AlitripBtripOpenplatformAddressGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripOpenplatformAddressGetRequest) SetRq(rq *OpenApiJumpInfoRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripOpenplatformAddressGetRequest) GetRq() *OpenApiJumpInfoRq {
    return r.rq
}

