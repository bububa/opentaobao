package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformCtsOrderStopAPIRequest CTS截断订单 API请求
// tmall.traceplatform.cts.order.stop
//
// 截断CTS订单
type TmallTraceplatformCtsOrderStopAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *TraceInfo
}

// NewTmallTraceplatformCtsOrderStopRequest 初始化TmallTraceplatformCtsOrderStopAPIRequest对象
func NewTmallTraceplatformCtsOrderStopRequest() *TmallTraceplatformCtsOrderStopAPIRequest {
	return &TmallTraceplatformCtsOrderStopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.cts.order.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformCtsOrderStopAPIRequest) SetTraceInfo(_traceInfo *TraceInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// Get TraceInfo Getter
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetTraceInfo() *TraceInfo {
	return r._traceInfo
}
