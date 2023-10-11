package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest 预约单差异明细查询 API请求
// cainiao.global.im.pickup.appointment.order.difference.detail
//
// 预约单差异明细查询
type CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest struct {
	model.Params
	// 请求参数
	_statusRequest *AppointmentOrderStatusRequest
}

// NewCainiaoGlobalImPickupAppointmentOrderDifferenceDetailRequest 初始化CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest对象
func NewCainiaoGlobalImPickupAppointmentOrderDifferenceDetailRequest() *CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest {
	return &CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.difference.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatusRequest is StatusRequest Setter
// 请求参数
func (r *CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest) SetStatusRequest(_statusRequest *AppointmentOrderStatusRequest) error {
	r._statusRequest = _statusRequest
	r.Set("status_request", _statusRequest)
	return nil
}

// GetStatusRequest StatusRequest Getter
func (r CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest) GetStatusRequest() *AppointmentOrderStatusRequest {
	return r._statusRequest
}
