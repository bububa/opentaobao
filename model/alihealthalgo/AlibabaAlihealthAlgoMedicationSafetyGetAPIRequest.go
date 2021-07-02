package alihealthalgo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest 合理用药api API请求
// alibaba.alihealth.algo.medication.safety.get
//
// 合理用药规则引擎服务
type AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest struct {
	model.Params
	// 业务请求对象
	_paramSolutionRequestTopSupport *SolutionRequestTopSupport
}

// NewAlibabaAlihealthAlgoMedicationSafetyGetRequest 初始化AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest对象
func NewAlibabaAlihealthAlgoMedicationSafetyGetRequest() *AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest {
	return &AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.algo.medication.safety.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamSolutionRequestTopSupport is ParamSolutionRequestTopSupport Setter
// 业务请求对象
func (r *AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) SetParamSolutionRequestTopSupport(_paramSolutionRequestTopSupport *SolutionRequestTopSupport) error {
	r._paramSolutionRequestTopSupport = _paramSolutionRequestTopSupport
	r.Set("param_solution_request_top_support", _paramSolutionRequestTopSupport)
	return nil
}

// GetParamSolutionRequestTopSupport ParamSolutionRequestTopSupport Getter
func (r AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) GetParamSolutionRequestTopSupport() *SolutionRequestTopSupport {
	return r._paramSolutionRequestTopSupport
}
