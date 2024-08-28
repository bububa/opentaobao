package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatUserLogin 微信小程序用户登录
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
func AlitripMerchantGalaxyWechatUserLogin(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatUserLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
