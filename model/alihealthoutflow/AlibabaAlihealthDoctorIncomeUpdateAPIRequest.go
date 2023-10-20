package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorIncomeUpdateAPIRequest 医蝶谷医生收入打款情况回调 API请求
// alibaba.alihealth.doctor.income.update
//
// 医蝶谷医生收入打款情况回调
type AlibabaAlihealthDoctorIncomeUpdateAPIRequest struct {
	model.Params
	// 入参
	_doctorIncomeStatusRequest *DoctorIncomeStatusRequest
}

// NewAlibabaAlihealthDoctorIncomeUpdateRequest 初始化AlibabaAlihealthDoctorIncomeUpdateAPIRequest对象
func NewAlibabaAlihealthDoctorIncomeUpdateRequest() *AlibabaAlihealthDoctorIncomeUpdateAPIRequest {
	return &AlibabaAlihealthDoctorIncomeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDoctorIncomeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.doctor.income.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDoctorIncomeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDoctorIncomeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDoctorIncomeStatusRequest is DoctorIncomeStatusRequest Setter
// 入参
func (r *AlibabaAlihealthDoctorIncomeUpdateAPIRequest) SetDoctorIncomeStatusRequest(_doctorIncomeStatusRequest *DoctorIncomeStatusRequest) error {
	r._doctorIncomeStatusRequest = _doctorIncomeStatusRequest
	r.Set("doctor_income_status_request", _doctorIncomeStatusRequest)
	return nil
}

// GetDoctorIncomeStatusRequest DoctorIncomeStatusRequest Getter
func (r AlibabaAlihealthDoctorIncomeUpdateAPIRequest) GetDoctorIncomeStatusRequest() *DoctorIncomeStatusRequest {
	return r._doctorIncomeStatusRequest
}
