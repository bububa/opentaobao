package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预约单核销接口 API请求
alibaba.alihealth.reservation.order.verify

预约单核销
*/
type AlibabaAlihealthReservationOrderVerifyAPIRequest struct {
    model.Params
    // 请求参数
    _verify   *VerifyOrderRequest
}

// 初始化AlibabaAlihealthReservationOrderVerifyAPIRequest对象
func NewAlibabaAlihealthReservationOrderVerifyRequest() *AlibabaAlihealthReservationOrderVerifyAPIRequest{
    return &AlibabaAlihealthReservationOrderVerifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReservationOrderVerifyAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reservation.order.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReservationOrderVerifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Verify Setter
// 请求参数
func (r *AlibabaAlihealthReservationOrderVerifyAPIRequest) SetVerify(_verify *VerifyOrderRequest) error {
    r._verify = _verify
    r.Set("verify", _verify)
    return nil
}

// Verify Getter
func (r AlibabaAlihealthReservationOrderVerifyAPIRequest) GetVerify() *VerifyOrderRequest {
    return r._verify
}
