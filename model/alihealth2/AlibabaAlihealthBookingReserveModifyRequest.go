package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改预约 APIRequest
alibaba.alihealth.booking.reserve.modify

消费医疗统一预约平台，取消预约
*/
type AlibabaAlihealthBookingReserveModifyRequest struct {
    model.Params

    // modify
    modify   *IsvReserveRequest 

}

func NewAlibabaAlihealthBookingReserveModifyRequest() *AlibabaAlihealthBookingReserveModifyRequest{
    return &AlibabaAlihealthBookingReserveModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBookingReserveModifyRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.modify"
}

func (r AlibabaAlihealthBookingReserveModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBookingReserveModifyRequest) SetModify(modify *IsvReserveRequest) error {
    r.modify = modify
    r.Set("modify", modify)
    return nil
}

func (r AlibabaAlihealthBookingReserveModifyRequest) GetModify() *IsvReserveRequest {
    return r.modify
}

