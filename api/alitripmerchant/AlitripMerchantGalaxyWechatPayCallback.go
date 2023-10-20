package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatpaycallback 微信支付回调
// alitrip.merchant.galaxy.wechat.pay.callback
//
// 微信支付回调
func Alitripmerchantgalaxywechatpaycallback(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatpaycallbackAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatpaycallbackAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatpaycallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
