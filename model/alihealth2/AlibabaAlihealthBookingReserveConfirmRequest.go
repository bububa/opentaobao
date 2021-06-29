package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认预约 API请求
alibaba.alihealth.booking.reserve.confirm

确认预约
*/
type AlibabaAlihealthBookingReserveConfirmRequest struct {
    model.Params
    // 参数
    confirm   *IsvReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveConfirmRequest对象
func NewAlibabaAlihealthBookingReserveConfirmRequest() *AlibabaAlihealthBookingReserveConfirmRequest{
    return &AlibabaAlihealthBookingReserveConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveConfirmRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Confirm Setter
// 参数
func (r *AlibabaAlihealthBookingReserveConfirmRequest) SetConfirm(confirm *IsvReserveRequest) error {
    r.confirm = confirm
    r.Set("confirm", confirm)
    return nil
}

// Confirm Getter
func (r AlibabaAlihealthBookingReserveConfirmRequest) GetConfirm() *IsvReserveRequest {
    return r.confirm
}
