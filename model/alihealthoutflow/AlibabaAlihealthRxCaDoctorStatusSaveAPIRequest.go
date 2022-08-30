package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest ca认证获取医师认证结果 API请求
// alibaba.alihealth.rx.ca.doctor.status.save
//
// ca认证获取医师认证结果
type AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest struct {
	model.Params
	// 入参实体
	_doctorStatusDto *DoctorStatusDto
}

// NewAlibabaAlihealthRxCaDoctorStatusSaveRequest 初始化AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest对象
func NewAlibabaAlihealthRxCaDoctorStatusSaveRequest() *AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest {
	return &AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.ca.doctor.status.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDoctorStatusDto is DoctorStatusDto Setter
// 入参实体
func (r *AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest) SetDoctorStatusDto(_doctorStatusDto *DoctorStatusDto) error {
	r._doctorStatusDto = _doctorStatusDto
	r.Set("doctor_status_dto", _doctorStatusDto)
	return nil
}

// GetDoctorStatusDto DoctorStatusDto Getter
func (r AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest) GetDoctorStatusDto() *DoctorStatusDto {
	return r._doctorStatusDto
}
