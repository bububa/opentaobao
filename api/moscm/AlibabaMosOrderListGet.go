package moscm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosOrderListGet 批量查询订单交易
// alibaba.mos.order.list.get
//
// 批量查询交易信息
func AlibabaMosOrderListGet(ctx context.Context, clt *core.SDKClient, req *moscm.AlibabaMosOrderListGetAPIRequest, resp *moscm.AlibabaMosOrderListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
