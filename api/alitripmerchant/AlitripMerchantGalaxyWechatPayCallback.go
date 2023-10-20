package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatPayCallback 微信支付回调
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
func AlitripMerchantGalaxyWechatPayCallback(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
