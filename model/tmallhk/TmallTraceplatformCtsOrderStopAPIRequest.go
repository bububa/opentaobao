package tmallhk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTraceplatformCtsOrderStopAPIRequest) Reset() {
	r._traceInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.cts.order.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceInfo is TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformCtsOrderStopAPIRequest) SetTraceInfo(_traceInfo *TraceInfo) error {
	r._traceInfo = _traceInfo
	r.Set("trace_info", _traceInfo)
	return nil
}

// GetTraceInfo TraceInfo Getter
func (r TmallTraceplatformCtsOrderStopAPIRequest) GetTraceInfo() *TraceInfo {
	return r._traceInfo
}

var poolTmallTraceplatformCtsOrderStopAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTraceplatformCtsOrderStopRequest()
	},
}

// GetTmallTraceplatformCtsOrderStopRequest 从 sync.Pool 获取 TmallTraceplatformCtsOrderStopAPIRequest
func GetTmallTraceplatformCtsOrderStopAPIRequest() *TmallTraceplatformCtsOrderStopAPIRequest {
	return poolTmallTraceplatformCtsOrderStopAPIRequest.Get().(*TmallTraceplatformCtsOrderStopAPIRequest)
}

// ReleaseTmallTraceplatformCtsOrderStopAPIRequest 将 TmallTraceplatformCtsOrderStopAPIRequest 放入 sync.Pool
func ReleaseTmallTraceplatformCtsOrderStopAPIRequest(v *TmallTraceplatformCtsOrderStopAPIRequest) {
	v.Reset()
	poolTmallTraceplatformCtsOrderStopAPIRequest.Put(v)
}
