package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeRender 渲染订单价格
// taobao.openmall.trade.render
//
// 请求渲染订单价格
func TaobaoOpenmallTradeRender(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeRenderAPIRequest, resp *openmall.TaobaoOpenmallTradeRenderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
