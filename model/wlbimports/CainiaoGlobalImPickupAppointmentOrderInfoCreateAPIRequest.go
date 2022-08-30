package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest 首公里揽收-预约单创建 API请求
// cainiao.global.im.pickup.appointment.order.info.create
//
// 预约单创建
type CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest struct {
	model.Params
	// 预约单创建入参
	_appointmentCreateRequest *AppointmentCreateRequest
}

// NewCainiaoGlobalImPickupAppointmentOrderInfoCreateRequest 初始化CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest对象
func NewCainiaoGlobalImPickupAppointmentOrderInfoCreateRequest() *CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest {
	return &CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.info.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAppointmentCreateRequest is AppointmentCreateRequest Setter
// 预约单创建入参
func (r *CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) SetAppointmentCreateRequest(_appointmentCreateRequest *AppointmentCreateRequest) error {
	r._appointmentCreateRequest = _appointmentCreateRequest
	r.Set("appointment_create_request", _appointmentCreateRequest)
	return nil
}

// GetAppointmentCreateRequest AppointmentCreateRequest Getter
func (r CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) GetAppointmentCreateRequest() *AppointmentCreateRequest {
	return r._appointmentCreateRequest
}
