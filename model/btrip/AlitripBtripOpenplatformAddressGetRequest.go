package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】开放平台对外页面跳转 API请求
alitrip.btrip.openplatform.address.get

获取类目预定页跳转地址
*/
type AlitripBtripOpenplatformAddressGetRequest struct {
    model.Params
    // 入参
    _rq   *OpenApiJumpInfoRq
}

// 初始化AlitripBtripOpenplatformAddressGetRequest对象
func NewAlitripBtripOpenplatformAddressGetRequest() *AlitripBtripOpenplatformAddressGetRequest{
    return &AlitripBtripOpenplatformAddressGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenplatformAddressGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.openplatform.address.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenplatformAddressGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripOpenplatformAddressGetRequest) SetRq(_rq *OpenApiJumpInfoRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripOpenplatformAddressGetRequest) GetRq() *OpenApiJumpInfoRq {
    return r._rq
}
