package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosflowWorkStartprocessAPIRequest 发起流程 API请求
// alibaba.mosflow.work.startprocess
//
// 业务发起流程审批
type AlibabaMosflowWorkStartprocessAPIRequest struct {
	model.Params
	// 参数二:额外业务参数,Map的JSON串
	_variables string
	// 流程必传参数
	_parameterEntity *ParameterEntity
}

// NewAlibabaMosflowWorkStartprocessRequest 初始化AlibabaMosflowWorkStartprocessAPIRequest对象
func NewAlibabaMosflowWorkStartprocessRequest() *AlibabaMosflowWorkStartprocessAPIRequest {
	return &AlibabaMosflowWorkStartprocessAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosflowWorkStartprocessAPIRequest) Reset() {
	r._variables = ""
	r._parameterEntity = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.mosflow.work.startprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVariables is Variables Setter
// 参数二:额外业务参数,Map的JSON串
func (r *AlibabaMosflowWorkStartprocessAPIRequest) SetVariables(_variables string) error {
	r._variables = _variables
	r.Set("variables", _variables)
	return nil
}

// GetVariables Variables Getter
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetVariables() string {
	return r._variables
}

// SetParameterEntity is ParameterEntity Setter
// 流程必传参数
func (r *AlibabaMosflowWorkStartprocessAPIRequest) SetParameterEntity(_parameterEntity *ParameterEntity) error {
	r._parameterEntity = _parameterEntity
	r.Set("parameter_entity", _parameterEntity)
	return nil
}

// GetParameterEntity ParameterEntity Getter
func (r AlibabaMosflowWorkStartprocessAPIRequest) GetParameterEntity() *ParameterEntity {
	return r._parameterEntity
}

var poolAlibabaMosflowWorkStartprocessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosflowWorkStartprocessRequest()
	},
}

// GetAlibabaMosflowWorkStartprocessRequest 从 sync.Pool 获取 AlibabaMosflowWorkStartprocessAPIRequest
func GetAlibabaMosflowWorkStartprocessAPIRequest() *AlibabaMosflowWorkStartprocessAPIRequest {
	return poolAlibabaMosflowWorkStartprocessAPIRequest.Get().(*AlibabaMosflowWorkStartprocessAPIRequest)
}

// ReleaseAlibabaMosflowWorkStartprocessAPIRequest 将 AlibabaMosflowWorkStartprocessAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosflowWorkStartprocessAPIRequest(v *AlibabaMosflowWorkStartprocessAPIRequest) {
	v.Reset()
	poolAlibabaMosflowWorkStartprocessAPIRequest.Put(v)
}
