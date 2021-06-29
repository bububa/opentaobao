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
type AlibabaAlihealthReservationOrderVerifyRequest struct {
    model.Params
    // 请求参数
    verify   *VerifyOrderRequest
}

// 初始化AlibabaAlihealthReservationOrderVerifyRequest对象
func NewAlibabaAlihealthReservationOrderVerifyRequest() *AlibabaAlihealthReservationOrderVerifyRequest{
    return &AlibabaAlihealthReservationOrderVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReservationOrderVerifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reservation.order.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReservationOrderVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Verify Setter
// 请求参数
func (r *AlibabaAlihealthReservationOrderVerifyRequest) SetVerify(verify *VerifyOrderRequest) error {
    r.verify = verify
    r.Set("verify", verify)
    return nil
}

// Verify Getter
func (r AlibabaAlihealthReservationOrderVerifyRequest) GetVerify() *VerifyOrderRequest {
    return r.verify
}
