package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV 新增/修改复诊预约信息 API请求
alibaba.alihealth.booking.reserve.rise

ISV 新增/修改复诊预约信息
*/
type AlibabaAlihealthBookingReserveRiseAPIRequest struct {
    model.Params
    // 参数
    _riseRequest   *IsvRiseReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveRiseAPIRequest对象
func NewAlibabaAlihealthBookingReserveRiseRequest() *AlibabaAlihealthBookingReserveRiseAPIRequest{
    return &AlibabaAlihealthBookingReserveRiseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.rise"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RiseRequest Setter
// 参数
func (r *AlibabaAlihealthBookingReserveRiseAPIRequest) SetRiseRequest(_riseRequest *IsvRiseReserveRequest) error {
    r._riseRequest = _riseRequest
    r.Set("rise_request", _riseRequest)
    return nil
}

// RiseRequest Getter
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetRiseRequest() *IsvRiseReserveRequest {
    return r._riseRequest
}
