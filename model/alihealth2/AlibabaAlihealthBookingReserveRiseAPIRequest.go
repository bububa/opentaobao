package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveRiseAPIRequest ISV 新增/修改复诊预约信息 API请求
// alibaba.alihealth.booking.reserve.rise
//
// ISV 新增/修改复诊预约信息
type AlibabaAlihealthBookingReserveRiseAPIRequest struct {
	model.Params
	// 参数
	_riseRequest *IsvRiseReserveRequest
}

// NewAlibabaAlihealthBookingReserveRiseRequest 初始化AlibabaAlihealthBookingReserveRiseAPIRequest对象
func NewAlibabaAlihealthBookingReserveRiseRequest() *AlibabaAlihealthBookingReserveRiseAPIRequest {
	return &AlibabaAlihealthBookingReserveRiseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.rise"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRiseRequest is RiseRequest Setter
// 参数
func (r *AlibabaAlihealthBookingReserveRiseAPIRequest) SetRiseRequest(_riseRequest *IsvRiseReserveRequest) error {
	r._riseRequest = _riseRequest
	r.Set("rise_request", _riseRequest)
	return nil
}

// GetRiseRequest RiseRequest Getter
func (r AlibabaAlihealthBookingReserveRiseAPIRequest) GetRiseRequest() *IsvRiseReserveRequest {
	return r._riseRequest
}
