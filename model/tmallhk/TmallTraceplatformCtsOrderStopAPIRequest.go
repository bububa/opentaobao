package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformctsorderstopAPIRequest CTS截断订单 API请求
// tmall.traceplatform.cts.order.stop
//
// 截断CTS订单
type TmalltraceplatformctsorderstopAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *TraceInfo
}

// NewTmalltraceplatformctsorderstopRequest 初始化TmalltraceplatformctsorderstopAPIRequest对象
func NewTmalltraceplatformctsorderstopRequest() *TmalltraceplatformctsorderstopAPIRequest {
	return &TmalltraceplatformctsorderstopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltraceplatformctsorderstopAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.cts.order.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltraceplatformctsorderstopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltraceplatformctsorderstopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceInfo is TraceInfo Setter
// 入参traceInfo
func (r *TmalltraceplatformctsorderstopAPIRequest) SetTraceInfo(_traceInfo *TraceInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// GetTraceInfo TraceInfo Getter
func (r TmalltraceplatformctsorderstopAPIRequest) GetTraceInfo() *TraceInfo {
	return r._traceInfo
}
