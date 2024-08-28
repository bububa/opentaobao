package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeConfirm 确认收货
// taobao.openmall.trade.confirm
//
// 确认订单收货
func TaobaoOpenmallTradeConfirm(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeConfirmAPIRequest, resp *openmall.TaobaoOpenmallTradeConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
