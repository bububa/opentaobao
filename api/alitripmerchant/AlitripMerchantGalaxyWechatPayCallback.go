package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatPayCallback 微信支付回调
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
func AlitripMerchantGalaxyWechatPayCallback(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatPayCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
