package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预约单核销接口 APIRequest
alibaba.alihealth.reservation.order.verify

预约单核销
*/
type AlibabaAlihealthReservationOrderVerifyRequest struct {
    model.Params

    // 请求参数
    verify   *VerifyOrderRequest 

}

func NewAlibabaAlihealthReservationOrderVerifyRequest() *AlibabaAlihealthReservationOrderVerifyRequest{
    return &AlibabaAlihealthReservationOrderVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthReservationOrderVerifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reservation.order.verify"
}

func (r AlibabaAlihealthReservationOrderVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthReservationOrderVerifyRequest) SetVerify(verify *VerifyOrderRequest) error {
    r.verify = verify
    r.Set("verify", verify)
    return nil
}

func (r AlibabaAlihealthReservationOrderVerifyRequest) GetVerify() *VerifyOrderRequest {
    return r.verify
}

