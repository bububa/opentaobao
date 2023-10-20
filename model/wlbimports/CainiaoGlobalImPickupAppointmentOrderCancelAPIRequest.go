package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupappointmentordercancelAPIRequest 首公里揽收-取消预约单 API请求
// cainiao.global.im.pickup.appointment.order.cancel
//
// 首公里揽收-取消预约单创建
type CainiaoglobalimpickupappointmentordercancelAPIRequest struct {
	model.Params
	// 预约单取消请求参数
	_appointmentOrderCancelRequset *AppointmentOrderCancelRequset
}

// NewCainiaoglobalimpickupappointmentordercancelRequest 初始化CainiaoglobalimpickupappointmentordercancelAPIRequest对象
func NewCainiaoglobalimpickupappointmentordercancelRequest() *CainiaoglobalimpickupappointmentordercancelAPIRequest {
	return &CainiaoglobalimpickupappointmentordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupappointmentordercancelAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupappointmentordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupappointmentordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppointmentOrderCancelRequset is AppointmentOrderCancelRequset Setter
// 预约单取消请求参数
func (r *CainiaoglobalimpickupappointmentordercancelAPIRequest) SetAppointmentOrderCancelRequset(_appointmentOrderCancelRequset *AppointmentOrderCancelRequset) error {
	r._appointmentOrderCancelRequset = _appointmentOrderCancelRequset
	r.Set("appointment_order_cancel_requset", _appointmentOrderCancelRequset)
	return nil
}

// GetAppointmentOrderCancelRequset AppointmentOrderCancelRequset Getter
func (r CainiaoglobalimpickupappointmentordercancelAPIRequest) GetAppointmentOrderCancelRequset() *AppointmentOrderCancelRequset {
	return r._appointmentOrderCancelRequset
}
