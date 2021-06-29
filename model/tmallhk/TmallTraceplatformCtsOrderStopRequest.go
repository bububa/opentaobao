package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CTS截断订单 API请求
tmall.traceplatform.cts.order.stop

截断CTS订单
*/
type TmallTraceplatformCtsOrderStopRequest struct {
    model.Params
    // 入参traceInfo
    traceInfo   *TraceInfo
}

// 初始化TmallTraceplatformCtsOrderStopRequest对象
func NewTmallTraceplatformCtsOrderStopRequest() *TmallTraceplatformCtsOrderStopRequest{
    return &TmallTraceplatformCtsOrderStopRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCtsOrderStopRequest) GetApiMethodName() string {
    return "tmall.traceplatform.cts.order.stop"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCtsOrderStopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformCtsOrderStopRequest) SetTraceInfo(traceInfo *TraceInfo) error {
    r.traceInfo = traceInfo
    r.Set("trace_info", traceInfo)
    return nil
}

// TraceInfo Getter
func (r TmallTraceplatformCtsOrderStopRequest) GetTraceInfo() *TraceInfo {
    return r.traceInfo
}
