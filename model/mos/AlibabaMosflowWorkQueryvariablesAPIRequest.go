package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosflowWorkQueryvariablesAPIRequest 获取指定流程上下文参数 API请求
// alibaba.mosflow.work.queryvariables
//
// 业务查询指定流程上下文内容
type AlibabaMosflowWorkQueryvariablesAPIRequest struct {
	model.Params
	// 流程实例ID
	_processInstanceId string
}

// NewAlibabaMosflowWorkQueryvariablesRequest 初始化AlibabaMosflowWorkQueryvariablesAPIRequest对象
func NewAlibabaMosflowWorkQueryvariablesRequest() *AlibabaMosflowWorkQueryvariablesAPIRequest {
	return &AlibabaMosflowWorkQueryvariablesAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosflowWorkQueryvariablesAPIRequest) Reset() {
	r._processInstanceId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosflowWorkQueryvariablesAPIRequest) GetApiMethodName() string {
	return "alibaba.mosflow.work.queryvariables"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosflowWorkQueryvariablesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosflowWorkQueryvariablesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProcessInstanceId is ProcessInstanceId Setter
// 流程实例ID
func (r *AlibabaMosflowWorkQueryvariablesAPIRequest) SetProcessInstanceId(_processInstanceId string) error {
	r._processInstanceId = _processInstanceId
	r.Set("process_instance_id", _processInstanceId)
	return nil
}

// GetProcessInstanceId ProcessInstanceId Getter
func (r AlibabaMosflowWorkQueryvariablesAPIRequest) GetProcessInstanceId() string {
	return r._processInstanceId
}

var poolAlibabaMosflowWorkQueryvariablesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosflowWorkQueryvariablesRequest()
	},
}

// GetAlibabaMosflowWorkQueryvariablesRequest 从 sync.Pool 获取 AlibabaMosflowWorkQueryvariablesAPIRequest
func GetAlibabaMosflowWorkQueryvariablesAPIRequest() *AlibabaMosflowWorkQueryvariablesAPIRequest {
	return poolAlibabaMosflowWorkQueryvariablesAPIRequest.Get().(*AlibabaMosflowWorkQueryvariablesAPIRequest)
}

// ReleaseAlibabaMosflowWorkQueryvariablesAPIRequest 将 AlibabaMosflowWorkQueryvariablesAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosflowWorkQueryvariablesAPIRequest(v *AlibabaMosflowWorkQueryvariablesAPIRequest) {
	v.Reset()
	poolAlibabaMosflowWorkQueryvariablesAPIRequest.Put(v)
}
