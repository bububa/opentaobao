package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeGet 查询订单详情
// taobao.openmall.trade.get
//
// 查询订单详情
func TaobaoOpenmallTradeGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeGetAPIRequest, resp *openmall.TaobaoOpenmallTradeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
