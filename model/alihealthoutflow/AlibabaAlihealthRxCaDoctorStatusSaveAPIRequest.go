package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthrxcadoctorstatussaveAPIRequest ca认证获取医师认证结果 API请求
// alibaba.alihealth.rx.ca.doctor.status.save
//
// ca认证获取医师认证结果
type AlibabaalihealthrxcadoctorstatussaveAPIRequest struct {
	model.Params
	// 入参实体
	_doctorStatusDto *DoctorStatusDto
}

// NewAlibabaalihealthrxcadoctorstatussaveRequest 初始化AlibabaalihealthrxcadoctorstatussaveAPIRequest对象
func NewAlibabaalihealthrxcadoctorstatussaveRequest() *AlibabaalihealthrxcadoctorstatussaveAPIRequest {
	return &AlibabaalihealthrxcadoctorstatussaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthrxcadoctorstatussaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.ca.doctor.status.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthrxcadoctorstatussaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthrxcadoctorstatussaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDoctorStatusDto is DoctorStatusDto Setter
// 入参实体
func (r *AlibabaalihealthrxcadoctorstatussaveAPIRequest) SetDoctorStatusDto(_doctorStatusDto *DoctorStatusDto) error {
	r._doctorStatusDto = _doctorStatusDto
	r.Set("doctor_status_dto", _doctorStatusDto)
	return nil
}

// GetDoctorStatusDto DoctorStatusDto Getter
func (r AlibabaalihealthrxcadoctorstatussaveAPIRequest) GetDoctorStatusDto() *DoctorStatusDto {
	return r._doctorStatusDto
}
