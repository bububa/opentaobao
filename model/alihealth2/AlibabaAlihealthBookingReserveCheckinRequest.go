package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认到店 API请求
alibaba.alihealth.booking.reserve.checkin

消费医疗统一预约平台，ISV 确认到店
*/
type AlibabaAlihealthBookingReserveCheckinAPIRequest struct {
    model.Params
    // check_in
    _checkIn   *IsvReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveCheckinAPIRequest对象
func NewAlibabaAlihealthBookingReserveCheckinRequest() *AlibabaAlihealthBookingReserveCheckinAPIRequest{
    return &AlibabaAlihealthBookingReserveCheckinAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.checkin"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CheckIn Setter
// check_in
func (r *AlibabaAlihealthBookingReserveCheckinAPIRequest) SetCheckIn(_checkIn *IsvReserveRequest) error {
    r._checkIn = _checkIn
    r.Set("check_in", _checkIn)
    return nil
}

// CheckIn Getter
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetCheckIn() *IsvReserveRequest {
    return r._checkIn
}
