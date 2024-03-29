package wlbimports

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) Reset() {
	r._appointmentOrderStatusRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoGlobalImPickupAppointmentOrderStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupAppointmentOrderStatusRequest()
	},
}

// GetCainiaoGlobalImPickupAppointmentOrderStatusRequest 从 sync.Pool 获取 CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest
func GetCainiaoGlobalImPickupAppointmentOrderStatusAPIRequest() *CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest {
	return poolCainiaoGlobalImPickupAppointmentOrderStatusAPIRequest.Get().(*CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest)
}

// ReleaseCainiaoGlobalImPickupAppointmentOrderStatusAPIRequest 将 CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupAppointmentOrderStatusAPIRequest(v *CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupAppointmentOrderStatusAPIRequest.Put(v)
}
