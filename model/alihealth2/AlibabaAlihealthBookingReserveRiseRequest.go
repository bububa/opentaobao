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
type AlibabaAlihealthBookingReserveRiseRequest struct {
    model.Params
    // 参数
    _riseRequest   *IsvRiseReserveRequest
}

// 初始化AlibabaAlihealthBookingReserveRiseRequest对象
func NewAlibabaAlihealthBookingReserveRiseRequest() *AlibabaAlihealthBookingReserveRiseRequest{
    return &AlibabaAlihealthBookingReserveRiseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveRiseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.booking.reserve.rise"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveRiseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RiseRequest Setter
// 参数
func (r *AlibabaAlihealthBookingReserveRiseRequest) SetRiseRequest(_riseRequest *IsvRiseReserveRequest) error {
    r._riseRequest = _riseRequest
    r.Set("rise_request", _riseRequest)
    return nil
}

// RiseRequest Getter
func (r AlibabaAlihealthBookingReserveRiseRequest) GetRiseRequest() *IsvRiseReserveRequest {
    return r._riseRequest
}
