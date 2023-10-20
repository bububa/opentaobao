package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyorderworkcallbackAPIRequest 仓配作业结果回传接口 API请求
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
type AlibabaaelophyorderworkcallbackAPIRequest struct {
	model.Params
	// 仓配作业结果回传请求
	_workCallbackRequest *WorkCallbackRequest
}

// NewAlibabaaelophyorderworkcallbackRequest 初始化AlibabaaelophyorderworkcallbackAPIRequest对象
func NewAlibabaaelophyorderworkcallbackRequest() *AlibabaaelophyorderworkcallbackAPIRequest {
	return &AlibabaaelophyorderworkcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaelophyorderworkcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.work.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaelophyorderworkcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaelophyorderworkcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkCallbackRequest is WorkCallbackRequest Setter
// 仓配作业结果回传请求
func (r *AlibabaaelophyorderworkcallbackAPIRequest) SetWorkCallbackRequest(_workCallbackRequest *WorkCallbackRequest) error {
	r._workCallbackRequest = _workCallbackRequest
	r.Set("work_callback_request", _workCallbackRequest)
	return nil
}

// GetWorkCallbackRequest WorkCallbackRequest Getter
func (r AlibabaaelophyorderworkcallbackAPIRequest) GetWorkCallbackRequest() *WorkCallbackRequest {
	return r._workCallbackRequest
}
