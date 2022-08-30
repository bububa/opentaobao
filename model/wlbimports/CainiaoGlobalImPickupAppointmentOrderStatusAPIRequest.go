package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest 预约单状态查询 API请求
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
type CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest struct {
	model.Params
	// 请求对象
	_appointmentOrderStatusRequest *AppointmentOrderStatusRequest
}

// NewCainiaoGlobalImPickupAppointmentOrderStatusRequest 初始化CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest对象
func NewCainiaoGlobalImPickupAppointmentOrderStatusRequest() *CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest {
	return &CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAppointmentOrderStatusRequest is AppointmentOrderStatusRequest Setter
// 请求对象
func (r *CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) SetAppointmentOrderStatusRequest(_appointmentOrderStatusRequest *AppointmentOrderStatusRequest) error {
	r._appointmentOrderStatusRequest = _appointmentOrderStatusRequest
	r.Set("appointment_order_status_request", _appointmentOrderStatusRequest)
	return nil
}

// GetAppointmentOrderStatusRequest AppointmentOrderStatusRequest Getter
func (r CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) GetAppointmentOrderStatusRequest() *AppointmentOrderStatusRequest {
	return r._appointmentOrderStatusRequest
}
