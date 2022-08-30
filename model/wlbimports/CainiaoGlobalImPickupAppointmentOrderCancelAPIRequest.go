package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest 首公里揽收-取消预约单 API请求
// cainiao.global.im.pickup.appointment.order.cancel
//
// 首公里揽收-取消预约单创建
type CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest struct {
	model.Params
	// 预约单取消请求参数
	_appointmentOrderCancelRequset *AppointmentOrderCancelRequset
}

// NewCainiaoGlobalImPickupAppointmentOrderCancelRequest 初始化CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest对象
func NewCainiaoGlobalImPickupAppointmentOrderCancelRequest() *CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest {
	return &CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAppointmentOrderCancelRequset is AppointmentOrderCancelRequset Setter
// 预约单取消请求参数
func (r *CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) SetAppointmentOrderCancelRequset(_appointmentOrderCancelRequset *AppointmentOrderCancelRequset) error {
	r._appointmentOrderCancelRequset = _appointmentOrderCancelRequset
	r.Set("appointment_order_cancel_requset", _appointmentOrderCancelRequset)
	return nil
}

// GetAppointmentOrderCancelRequset AppointmentOrderCancelRequset Getter
func (r CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) GetAppointmentOrderCancelRequset() *AppointmentOrderCancelRequset {
	return r._appointmentOrderCancelRequset
}
