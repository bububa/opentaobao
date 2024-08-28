package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatLogin 星河-用户使用微信登陆
// alitrip.merchant.galaxy.wechat.login
//
// 星河产品=用户微信小程序登陆
func AlitripMerchantGalaxyWechatLogin(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatLoginAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
