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
type AlibabaAlihealthBookingReserveTreatAPIRequest struct {
    model.Params
    // treat
    _treat   *IsvReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveTreatAPIRequest对象
func NewAlibabaAlihealthBookingReserveTreatRequest() *AlibabaAlihealthBookingReserveTreatAPIRequest{
    return &AlibabaAlihealthBookingReserveTreatAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.treat"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Treat Setter
// treat
func (r *AlibabaAlihealthBookingReserveTreatAPIRequest) SetTreat(_treat *IsvReserveRequest) error {
    r._treat = _treat
    r.Set("treat", _treat)
    return nil
}

// Treat Getter
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetTreat() *IsvReserveRequest {
    return r._treat
}
