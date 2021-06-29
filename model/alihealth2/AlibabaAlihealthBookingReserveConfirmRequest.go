package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认预约 APIRequest
alibaba.alihealth.booking.reserve.confirm

确认预约
*/
type AlibabaAlihealthBookingReserveConfirmRequest struct {
    model.Params

    // 参数
    confirm   *IsvReserveRequest 

}

func NewAlibabaAlihealthBookingReserveConfirmRequest() *AlibabaAlihealthBookingReserveConfirmRequest{
    return &AlibabaAlihealthBookingReserveConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBookingReserveConfirmRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.confirm"
}

func (r AlibabaAlihealthBookingReserveConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBookingReserveConfirmRequest) SetConfirm(confirm *IsvReserveRequest) error {
    r.confirm = confirm
    r.Set("confirm", confirm)
    return nil
}

func (r AlibabaAlihealthBookingReserveConfirmRequest) GetConfirm() *IsvReserveRequest {
    return r.confirm
}

