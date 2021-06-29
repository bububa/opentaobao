package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认就诊 API请求
alibaba.alihealth.booking.reserve.treat

统一预约平台，ISV确认就诊
*/
type AlibabaAlihealthBookingReserveTreatRequest struct {
    model.Params
    // treat
    treat   *IsvReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveTreatRequest对象
func NewAlibabaAlihealthBookingReserveTreatRequest() *AlibabaAlihealthBookingReserveTreatRequest{
    return &AlibabaAlihealthBookingReserveTreatRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveTreatRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.treat"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveTreatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Treat Setter
// treat
func (r *AlibabaAlihealthBookingReserveTreatRequest) SetTreat(treat *IsvReserveRequest) error {
    r.treat = treat
    r.Set("treat", treat)
    return nil
}

// Treat Getter
func (r AlibabaAlihealthBookingReserveTreatRequest) GetTreat() *IsvReserveRequest {
    return r.treat
}
