package tmallhk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallhk"
)

/* 
CTS截断订单 
tmall.traceplatform.cts.order.stop

截断CTS订单
*/
func TmallTraceplatformCtsOrderStop(clt *core.SDKClient, req *tmallhk.TmallTraceplatformCtsOrderStopRequest, session string) (*tmallhk.TmallTraceplatformCtsOrderStopAPIResponse, error) {
    var resp tmallhk.TmallTraceplatformCtsOrderStopAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
