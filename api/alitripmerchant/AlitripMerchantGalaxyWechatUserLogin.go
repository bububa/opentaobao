package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatUserLogin 微信小程序用户登录
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
func AlitripMerchantGalaxyWechatUserLogin(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
