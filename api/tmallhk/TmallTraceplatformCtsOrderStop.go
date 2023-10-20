package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformCtsOrderStop CTS截断订单
// tmall.traceplatform.cts.order.stop
//
// 截断CTS订单
func TmallTraceplatformCtsOrderStop(clt *core.SDKClient, req *tmallhk.TmallTraceplatformCtsOrderStopAPIRequest, resp *tmallhk.TmallTraceplatformCtsOrderStopAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
