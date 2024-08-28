package tvpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayOrderPartnerpay tv支付第三方支付订单
// taobao.tvpay.order.partnerpay
//
// tv支付第三方发起并支付订单（使用设备授权）
func TaobaoTvpayOrderPartnerpay(ctx context.Context, clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderPartnerpayAPIRequest, resp *tvpay.TaobaoTvpayOrderPartnerpayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
