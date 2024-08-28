package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderShow 闲鱼回收订单查询V1.1
// alibaba.idle.recycle.order.show
//
// 查询回收订单
func AlibabaIdleRecycleOrderShow(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderShowAPIRequest, resp *idle.AlibabaIdleRecycleOrderShowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
