package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscOrderOrderUploadAPIRequest 订单回流 API请求
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
type AlibabaAlscOrderOrderUploadAPIRequest struct {
	model.Params
	// 订单回流参数
	_paramBackflowRequest *BackflowRequest
}

// NewAlibabaAlscOrderOrderUploadRequest 初始化AlibabaAlscOrderOrderUploadAPIRequest对象
func NewAlibabaAlscOrderOrderUploadRequest() *AlibabaAlscOrderOrderUploadAPIRequest {
	return &AlibabaAlscOrderOrderUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscOrderOrderUploadAPIRequest) Reset() {
	r._paramBackflowRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.order.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBackflowRequest is ParamBackflowRequest Setter
// 订单回流参数
func (r *AlibabaAlscOrderOrderUploadAPIRequest) SetParamBackflowRequest(_paramBackflowRequest *BackflowRequest) error {
	r._paramBackflowRequest = _paramBackflowRequest
	r.Set("param_backflow_request", _paramBackflowRequest)
	return nil
}

// GetParamBackflowRequest ParamBackflowRequest Getter
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetParamBackflowRequest() *BackflowRequest {
	return r._paramBackflowRequest
}

var poolAlibabaAlscOrderOrderUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscOrderOrderUploadRequest()
	},
}

// GetAlibabaAlscOrderOrderUploadRequest 从 sync.Pool 获取 AlibabaAlscOrderOrderUploadAPIRequest
func GetAlibabaAlscOrderOrderUploadAPIRequest() *AlibabaAlscOrderOrderUploadAPIRequest {
	return poolAlibabaAlscOrderOrderUploadAPIRequest.Get().(*AlibabaAlscOrderOrderUploadAPIRequest)
}

// ReleaseAlibabaAlscOrderOrderUploadAPIRequest 将 AlibabaAlscOrderOrderUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscOrderOrderUploadAPIRequest(v *AlibabaAlscOrderOrderUploadAPIRequest) {
	v.Reset()
	poolAlibabaAlscOrderOrderUploadAPIRequest.Put(v)
}
