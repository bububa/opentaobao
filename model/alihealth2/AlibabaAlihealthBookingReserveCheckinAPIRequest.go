package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbookingreservecheckinAPIRequest 确认到店 API请求
// alibaba.alihealth.booking.reserve.checkin
//
// 消费医疗统一预约平台，ISV 确认到店
type AlibabaalihealthbookingreservecheckinAPIRequest struct {
	model.Params
	// check_in
	_checkIn *IsvReserveRequest
}

// NewAlibabaalihealthbookingreservecheckinRequest 初始化AlibabaalihealthbookingreservecheckinAPIRequest对象
func NewAlibabaalihealthbookingreservecheckinRequest() *AlibabaalihealthbookingreservecheckinAPIRequest {
	return &AlibabaalihealthbookingreservecheckinAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbookingreservecheckinAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.checkin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbookingreservecheckinAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbookingreservecheckinAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckIn is CheckIn Setter
// check_in
func (r *AlibabaalihealthbookingreservecheckinAPIRequest) SetCheckIn(_checkIn *IsvReserveRequest) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r AlibabaalihealthbookingreservecheckinAPIRequest) GetCheckIn() *IsvReserveRequest {
	return r._checkIn
}
