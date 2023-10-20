package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderCodeExchange 核销券码
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
func TmallAlihouseTradeCouponOrderCodeExchange(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest, resp *alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
