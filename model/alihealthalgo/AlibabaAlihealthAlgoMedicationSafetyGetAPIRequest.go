package alihealthalgo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthalgomedicationsafetygetAPIRequest 合理用药api API请求
// alibaba.alihealth.algo.medication.safety.get
//
// 合理用药规则引擎服务
type AlibabaalihealthalgomedicationsafetygetAPIRequest struct {
	model.Params
	// 业务请求对象
	_paramSolutionRequestTopSupport *SolutionRequestTopSupport
}

// NewAlibabaalihealthalgomedicationsafetygetRequest 初始化AlibabaalihealthalgomedicationsafetygetAPIRequest对象
func NewAlibabaalihealthalgomedicationsafetygetRequest() *AlibabaalihealthalgomedicationsafetygetAPIRequest {
	return &AlibabaalihealthalgomedicationsafetygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthalgomedicationsafetygetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.algo.medication.safety.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthalgomedicationsafetygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthalgomedicationsafetygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSolutionRequestTopSupport is ParamSolutionRequestTopSupport Setter
// 业务请求对象
func (r *AlibabaalihealthalgomedicationsafetygetAPIRequest) SetParamSolutionRequestTopSupport(_paramSolutionRequestTopSupport *SolutionRequestTopSupport) error {
	r._paramSolutionRequestTopSupport = _paramSolutionRequestTopSupport
	r.Set("param_solution_request_top_support", _paramSolutionRequestTopSupport)
	return nil
}

// GetParamSolutionRequestTopSupport ParamSolutionRequestTopSupport Getter
func (r AlibabaalihealthalgomedicationsafetygetAPIRequest) GetParamSolutionRequestTopSupport() *SolutionRequestTopSupport {
	return r._paramSolutionRequestTopSupport
}
