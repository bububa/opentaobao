package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认就诊 APIRequest
alibaba.alihealth.booking.reserve.treat

统一预约平台，ISV确认就诊
*/
type AlibabaAlihealthBookingReserveTreatRequest struct {
    model.Params

    // treat
    treat   *IsvReserveRequest 

}

func NewAlibabaAlihealthBookingReserveTreatRequest() *AlibabaAlihealthBookingReserveTreatRequest{
    return &AlibabaAlihealthBookingReserveTreatRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBookingReserveTreatRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.treat"
}

func (r AlibabaAlihealthBookingReserveTreatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBookingReserveTreatRequest) SetTreat(treat *IsvReserveRequest) error {
    r.treat = treat
    r.Set("treat", treat)
    return nil
}

func (r AlibabaAlihealthBookingReserveTreatRequest) GetTreat() *IsvReserveRequest {
    return r.treat
}

