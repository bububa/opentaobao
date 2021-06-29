package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消预约 APIRequest
alibaba.alihealth.booking.reserve.cancel

消费医疗统一预约平台，ISV取消预约
*/
type AlibabaAlihealthBookingReserveCancelRequest struct {
    model.Params

    // cancel
    cancel   *IsvReserveRequest 

}

func NewAlibabaAlihealthBookingReserveCancelRequest() *AlibabaAlihealthBookingReserveCancelRequest{
    return &AlibabaAlihealthBookingReserveCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBookingReserveCancelRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.cancel"
}

func (r AlibabaAlihealthBookingReserveCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBookingReserveCancelRequest) SetCancel(cancel *IsvReserveRequest) error {
    r.cancel = cancel
    r.Set("cancel", cancel)
    return nil
}

func (r AlibabaAlihealthBookingReserveCancelRequest) GetCancel() *IsvReserveRequest {
    return r.cancel
}

