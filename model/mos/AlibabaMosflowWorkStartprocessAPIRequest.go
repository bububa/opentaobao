package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosflowworkstartprocessAPIRequest 发起流程 API请求
// alibaba.mosflow.work.startprocess
//
// 业务发起流程审批
type AlibabamosflowworkstartprocessAPIRequest struct {
	model.Params
	// 参数二:额外业务参数,Map的JSON串
	_variables string
	// 流程必传参数
	_parameterEntity *ParameterEntity
}

// NewAlibabamosflowworkstartprocessRequest 初始化AlibabamosflowworkstartprocessAPIRequest对象
func NewAlibabamosflowworkstartprocessRequest() *AlibabamosflowworkstartprocessAPIRequest {
	return &AlibabamosflowworkstartprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosflowworkstartprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.mosflow.work.startprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosflowworkstartprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosflowworkstartprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVariables is Variables Setter
// 参数二:额外业务参数,Map的JSON串
func (r *AlibabamosflowworkstartprocessAPIRequest) SetVariables(_variables string) error {
	r._variables = _variables
	r.Set("variables", _variables)
	return nil
}

// GetVariables Variables Getter
func (r AlibabamosflowworkstartprocessAPIRequest) GetVariables() string {
	return r._variables
}

// SetParameterEntity is ParameterEntity Setter
// 流程必传参数
func (r *AlibabamosflowworkstartprocessAPIRequest) SetParameterEntity(_parameterEntity *ParameterEntity) error {
	r._parameterEntity = _parameterEntity
	r.Set("parameter_entity", _parameterEntity)
	return nil
}

// GetParameterEntity ParameterEntity Getter
func (r AlibabamosflowworkstartprocessAPIRequest) GetParameterEntity() *ParameterEntity {
	return r._parameterEntity
}
