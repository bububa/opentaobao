package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeClose 关闭订单
// taobao.openmall.trade.close
//
// 关闭订单
func TaobaoOpenmallTradeClose(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeCloseAPIRequest, resp *openmall.TaobaoOpenmallTradeCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
