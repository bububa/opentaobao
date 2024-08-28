package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderCodeExchange 核销券码
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
func TmallAlihouseTradeCouponOrderCodeExchange(ctx context.Context, clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest, resp *alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
