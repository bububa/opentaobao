package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTraceplatformCtsOrderStopAPIRequest
CTS截断订单 API请求
tmall.traceplatform.cts.order.stop

截断CTS订单 */
type TmallTraceplatformCtsOrderStopAPIRequest struct {
	model.Params
	// 入参traceInfo
	_traceInfo *TraceInfo
}

// New
