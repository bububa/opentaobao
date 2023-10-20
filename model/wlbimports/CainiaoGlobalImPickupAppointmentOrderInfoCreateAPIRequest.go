package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupappointmentorderinfocreateAPIRequest 首公里揽收-预约单创建 API请求
// cainiao.global.im.pickup.appointment.order.info.create
//
// 预约单创建
type CainiaoglobalimpickupappointmentorderinfocreateAPIRequest struct {
	model.Params
	// 预约单创建入参
	_appointmentCreateRequest *AppointmentCreateRequest
}

// NewCainiaoglobalimpickupappointmentorderinfocreateRequest 初始化CainiaoglobalimpickupappointmentorderinfocreateAPIRequest对象
func NewCainiaoglobalimpickupappointmentorderinfocreateRequest() *CainiaoglobalimpickupappointmentorderinfocreateAPIRequest {
	return &CainiaoglobalimpickupappointmentorderinfocreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupappointmentorderinfocreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.appointment.order.info.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupappointmentorderinfocreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupappointmentorderinfocreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppointmentCreateRequest is AppointmentCreateRequest Setter
// 预约单创建入参
func (r *CainiaoglobalimpickupappointmentorderinfocreateAPIRequest) SetAppointmentCreateRequest(_appointmentCreateRequest *AppointmentCreateRequest) error {
	r._appointmentCreateRequest = _appointmentCreateRequest
	r.Set("appointment_create_request", _appointmentCreateRequest)
	return nil
}

// GetAppointmentCreateRequest AppointmentCreateRequest Getter
func (r CainiaoglobalimpickupappointmentorderinfocreateAPIRequest) GetAppointmentCreateRequest() *AppointmentCreateRequest {
	return r._appointmentCreateRequest
}
