package moscm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosOrderQuery 批量查询订单信息
// alibaba.mos.order.query
//
// 查询多笔交易信息
func AlibabaMosOrderQuery(ctx context.Context, clt *core.SDKClient, req *moscm.AlibabaMosOrderQueryAPIRequest, resp *moscm.AlibabaMosOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
