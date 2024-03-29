package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderWorkCallbackAPIRequest 仓配作业结果回传接口 API请求
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
type AlibabaAelophyOrderWorkCallbackAPIRequest struct {
	model.Params
	// 仓配作业结果回传请求
	_workCallbackRequest *WorkCallbackRequest
}

// NewAlibabaAelophyOrderWorkCallbackRequest 初始化AlibabaAelophyOrderWorkCallbackAPIRequest对象
func NewAlibabaAelophyOrderWorkCallbackRequest() *AlibabaAelophyOrderWorkCallbackAPIRequest {
	return &AlibabaAelophyOrderWorkCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAelophyOrderWorkCallbackAPIRequest) Reset() {
	r._workCallbackRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyOrderWorkCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.work.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyOrderWorkCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyOrderWorkCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkCallbackRequest is WorkCallbackRequest Setter
// 仓配作业结果回传请求
func (r *AlibabaAelophyOrderWorkCallbackAPIRequest) SetWorkCallbackRequest(_workCallbackRequest *WorkCallbackRequest) error {
	r._workCallbackRequest = _workCallbackRequest
	r.Set("work_callback_request", _workCallbackRequest)
	return nil
}

// GetWorkCallbackRequest WorkCallbackRequest Getter
func (r AlibabaAelophyOrderWorkCallbackAPIRequest) GetWorkCallbackRequest() *WorkCallbackRequest {
	return r._workCallbackRequest
}

var poolAlibabaAelophyOrderWorkCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAelophyOrderWorkCallbackRequest()
	},
}

// GetAlibabaAelophyOrderWorkCallbackRequest 从 sync.Pool 获取 AlibabaAelophyOrderWorkCallbackAPIRequest
func GetAlibabaAelophyOrderWorkCallbackAPIRequest() *AlibabaAelophyOrderWorkCallbackAPIRequest {
	return poolAlibabaAelophyOrderWorkCallbackAPIRequest.Get().(*AlibabaAelophyOrderWorkCallbackAPIRequest)
}

// ReleaseAlibabaAelophyOrderWorkCallbackAPIRequest 将 AlibabaAelophyOrderWorkCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAelophyOrderWorkCallbackAPIRequest(v *AlibabaAelophyOrderWorkCallbackAPIRequest) {
	v.Reset()
	poolAlibabaAelophyOrderWorkCallbackAPIRequest.Put(v)
}
