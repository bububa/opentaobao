package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatUserLogin 微信小程序用户登录
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
func AlitripMerchantGalaxyWechatUserLogin(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
