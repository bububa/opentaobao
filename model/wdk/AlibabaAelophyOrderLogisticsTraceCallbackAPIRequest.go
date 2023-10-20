package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest 配送轨迹回传 API请求
// alibaba.aelophy.order.logistics.trace.callback
//
// 配送轨迹回传
type AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest struct {
	model.Params
	// 配送轨迹回传请求
	_logisticsTraceCallbackRequest *LogisticsTraceCallbackRequest
}

// NewAlibabaAelophyOrderLogisticsTraceCallbackRequest 初始化AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest对象
func NewAlibabaAelophyOrderLogisticsTraceCallbackRequest() *AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest {
	return &AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) Reset() {
	r._logisticsTraceCallbackRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.logistics.trace.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsTraceCallbackRequest is LogisticsTraceCallbackRequest Setter
// 配送轨迹回传请求
func (r *AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) SetLogisticsTraceCallbackRequest(_logisticsTraceCallbackRequest *LogisticsTraceCallbackRequest) error {
	r._logisticsTraceCallbackRequest = _logisticsTraceCallbackRequest
	r.Set("logistics_trace_callback_request", _logisticsTraceCallbackRequest)
	return nil
}

// GetLogisticsTraceCallbackRequest LogisticsTraceCallbackRequest Getter
func (r AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) GetLogisticsTraceCallbackRequest() *LogisticsTraceCallbackRequest {
	return r._logisticsTraceCallbackRequest
}

var poolAlibabaAelophyOrderLogisticsTraceCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAelophyOrderLogisticsTraceCallbackRequest()
	},
}

// GetAlibabaAelophyOrderLogisticsTraceCallbackRequest 从 sync.Pool 获取 AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest
func GetAlibabaAelophyOrderLogisticsTraceCallbackAPIRequest() *AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest {
	return poolAlibabaAelophyOrderLogisticsTraceCallbackAPIRequest.Get().(*AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest)
}

// ReleaseAlibabaAelophyOrderLogisticsTraceCallbackAPIRequest 将 AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAelophyOrderLogisticsTraceCallbackAPIRequest(v *AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest) {
	v.Reset()
	poolAlibabaAelophyOrderLogisticsTraceCallbackAPIRequest.Put(v)
}
