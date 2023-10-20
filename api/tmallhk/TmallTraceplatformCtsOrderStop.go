package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmalltraceplatformctsorderstop CTS截断订单
// tmall.traceplatform.cts.order.stop
//
// 截断CTS订单
func Tmalltraceplatformctsorderstop(clt *core.SDKClient, req *tmallhk.TmalltraceplatformctsorderstopAPIRequest, session string) (*tmallhk.TmalltraceplatformctsorderstopAPIResponse, error) {
	var resp tmallhk.TmalltraceplatformctsorderstopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
