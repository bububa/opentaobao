package wlbimports

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) Reset() {
	r._appointmentCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.info.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupAppointmentOrderInfoCreateRequest()
	},
}

// GetCainiaoGlobalImPickupAppointmentOrderInfoCreateRequest 从 sync.Pool 获取 CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest
func GetCainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest() *CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest {
	return poolCainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest.Get().(*CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest)
}

// ReleaseCainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest 将 CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest(v *CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest.Put(v)
}
