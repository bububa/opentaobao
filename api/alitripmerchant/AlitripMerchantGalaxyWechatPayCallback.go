package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatPayCallback 微信支付回调
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
func AlitripMerchantGalaxyWechatPayCallback(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
