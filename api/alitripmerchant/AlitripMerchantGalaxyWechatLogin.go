package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

/* AlitripMerchantGalaxyWechatLogin
星河-用户使用微信登陆
alitrip.merchant.galaxy.wechat.login

星河产品=用户微信小程序登陆 */
func AlitripMerchantGalaxyWechatLogin(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatLoginAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatLoginAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatLoginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
