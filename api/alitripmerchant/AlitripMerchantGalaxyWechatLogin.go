package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatlogin 星河-用户使用微信登陆
// alitrip.merchant.galaxy.wechat.login
//
// 星河产品=用户微信小程序登陆
func Alitripmerchantgalaxywechatlogin(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatloginAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatloginAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
