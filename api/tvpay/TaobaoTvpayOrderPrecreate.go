package tvpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayOrderPrecreate tv支付预下单
// taobao.tvpay.order.precreate
//
// tv支付预下单
func TaobaoTvpayOrderPrecreate(ctx context.Context, clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderPrecreateAPIRequest, resp *tvpay.TaobaoTvpayOrderPrecreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
