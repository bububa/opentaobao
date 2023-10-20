package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupappointmentorderstatusAPIRequest 预约单状态查询 API请求
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
type CainiaoglobalimpickupappointmentorderstatusAPIRequest struct {
	model.Params
	// 请求对象
	_appointmentOrderStatusRequest *AppointmentOrderStatusRequest
}

// NewCainiaoglobalimpickupappointmentorderstatusRequest 初始化CainiaoglobalimpickupappointmentorderstatusAPIRequest对象
func NewCainiaoglobalimpickupappointmentorderstatusRequest() *CainiaoglobalimpickupappointmentorderstatusAPIRequest {
	return &CainiaoglobalimpickupappointmentorderstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupappointmentorderstatusAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupappointmentorderstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupappointmentorderstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppointmentOrderStatusRequest is AppointmentOrderStatusRequest Setter
// 请求对象
func (r *CainiaoglobalimpickupappointmentorderstatusAPIRequest) SetAppointmentOrderStatusRequest(_appointmentOrderStatusRequest *AppointmentOrderStatusRequest) error {
	r._appointmentOrderStatusRequest = _appointmentOrderStatusRequest
	r.Set("appointment_order_status_request", _appointmentOrderStatusRequest)
	return nil
}

// GetAppointmentOrderStatusRequest AppointmentOrderStatusRequest Getter
func (r CainiaoglobalimpickupappointmentorderstatusAPIRequest) GetAppointmentOrderStatusRequest() *AppointmentOrderStatusRequest {
	return r._appointmentOrderStatusRequest
}
