package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CTS截断订单 APIRequest
tmall.traceplatform.cts.order.stop

截断CTS订单
*/
type TmallTraceplatformCtsOrderStopRequest struct {
    model.Params

    // 入参traceInfo
    traceInfo   *TraceInfo 

}

func NewTmallTraceplatformCtsOrderStopRequest() *TmallTraceplatformCtsOrderStopRequest{
    return &TmallTraceplatformCtsOrderStopRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraceplatformCtsOrderStopRequest) GetApiMethodName() string {
    return "tmall.traceplatform.cts.order.stop"
}

func (r TmallTraceplatformCtsOrderStopRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraceplatformCtsOrderStopRequest) SetTraceInfo(traceInfo *TraceInfo) error {
    r.traceInfo = traceInfo
    r.Set("trace_info", traceInfo)
    return nil
}

func (r TmallTraceplatformCtsOrderStopRequest) GetTraceInfo() *TraceInfo {
    return r.traceInfo
}

