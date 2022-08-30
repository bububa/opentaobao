package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderCodeExchange 核销券码
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
func TmallAlihouseTradeCouponOrderCodeExchange(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest, session string) (*alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIResponse, error) {
	var resp alihouse.TmallAlihouseTradeCouponOrderCodeExchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
