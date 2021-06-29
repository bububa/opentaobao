package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV 新增/修改复诊预约信息 APIRequest
alibaba.alihealth.booking.reserve.rise

ISV 新增/修改复诊预约信息
*/
type AlibabaAlihealthBookingReserveRiseRequest struct {
    model.Params

    // 参数
    riseRequest   *IsvRiseReserveRequest 

}

func NewAlibabaAlihealthBookingReserveRiseRequest() *AlibabaAlihealthBookingReserveRiseRequest{
    return &AlibabaAlihealthBookingReserveRiseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBookingReserveRiseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.rise"
}

func (r AlibabaAlihealthBookingReserveRiseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBookingReserveRiseRequest) SetRiseRequest(riseRequest *IsvRiseReserveRequest) error {
    r.riseRequest = riseRequest
    r.Set("rise_request", riseRequest)
    return nil
}

func (r AlibabaAlihealthBookingReserveRiseRequest) GetRiseRequest() *IsvRiseReserveRequest {
    return r.riseRequest
}

