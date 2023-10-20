package wlbimports

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) Reset() {
	r._appointmentOrderCancelRequset = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoGlobalImPickupAppointmentOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupAppointmentOrderCancelRequest()
	},
}

// GetCainiaoGlobalImPickupAppointmentOrderCancelRequest 从 sync.Pool 获取 CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest
func GetCainiaoGlobalImPickupAppointmentOrderCancelAPIRequest() *CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest {
	return poolCainiaoGlobalImPickupAppointmentOrderCancelAPIRequest.Get().(*CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest)
}

// ReleaseCainiaoGlobalImPickupAppointmentOrderCancelAPIRequest 将 CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupAppointmentOrderCancelAPIRequest(v *CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupAppointmentOrderCancelAPIRequest.Put(v)
}
