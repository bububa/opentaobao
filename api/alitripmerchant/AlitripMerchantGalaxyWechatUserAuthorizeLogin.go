package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatUserAuthorizeLogin DFC-ID用户手机号授权登录
// alitrip.merchant.galaxy.wechat.user.authorize.login
//
// DFC-ID用户手机号授权登录
func AlitripMerchantGalaxyWechatUserAuthorizeLogin(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
