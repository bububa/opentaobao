package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosflowworkqueryvariablesAPIRequest 获取指定流程上下文参数 API请求
// alibaba.mosflow.work.queryvariables
//
// 业务查询指定流程上下文内容
type AlibabamosflowworkqueryvariablesAPIRequest struct {
	model.Params
	// 流程实例ID
	_processInstanceId string
}

// NewAlibabamosflowworkqueryvariablesRequest 初始化AlibabamosflowworkqueryvariablesAPIRequest对象
func NewAlibabamosflowworkqueryvariablesRequest() *AlibabamosflowworkqueryvariablesAPIRequest {
	return &AlibabamosflowworkqueryvariablesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosflowworkqueryvariablesAPIRequest) GetApiMethodName() string {
	return "alibaba.mosflow.work.queryvariables"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosflowworkqueryvariablesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosflowworkqueryvariablesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProcessInstanceId is ProcessInstanceId Setter
// 流程实例ID
func (r *AlibabamosflowworkqueryvariablesAPIRequest) SetProcessInstanceId(_processInstanceId string) error {
	r._processInstanceId = _processInstanceId
	r.Set("process_instance_id", _processInstanceId)
	return nil
}

// GetProcessInstanceId ProcessInstanceId Getter
func (r AlibabamosflowworkqueryvariablesAPIRequest) GetProcessInstanceId() string {
	return r._processInstanceId
}
