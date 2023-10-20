package alihealthalgo

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) Reset() {
	r._paramSolutionRequestTopSupport = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.algo.medication.safety.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthAlgoMedicationSafetyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthAlgoMedicationSafetyGetRequest()
	},
}

// GetAlibabaAlihealthAlgoMedicationSafetyGetRequest 从 sync.Pool 获取 AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest
func GetAlibabaAlihealthAlgoMedicationSafetyGetAPIRequest() *AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest {
	return poolAlibabaAlihealthAlgoMedicationSafetyGetAPIRequest.Get().(*AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest)
}

// ReleaseAlibabaAlihealthAlgoMedicationSafetyGetAPIRequest 将 AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthAlgoMedicationSafetyGetAPIRequest(v *AlibabaAlihealthAlgoMedicationSafetyGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthAlgoMedicationSafetyGetAPIRequest.Put(v)
}
