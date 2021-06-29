package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消预约 API请求
alibaba.alihealth.booking.reserve.cancel

消费医疗统一预约平台，ISV取消预约
*/
type AlibabaAlihealthBookingReserveCancelRequest struct {
    model.Params
    // cancel
    cancel   *IsvReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveCancelRequest对象
func NewAlibabaAlihealthBookingReserveCancelRequest() *AlibabaAlihealthBookingReserveCancelRequest{
    return &AlibabaAlihealthBookingReserveCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveCancelRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cancel Setter
// cancel
func (r *AlibabaAlihealthBookingReserveCancelRequest) SetCancel(cancel *IsvReserveRequest) error {
    r.cancel = cancel
    r.Set("cancel", cancel)
    return nil
}

// Cancel Getter
func (r AlibabaAlihealthBookingReserveCancelRequest) GetCancel() *IsvReserveRequest {
    return r.cancel
}
