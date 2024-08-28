package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatInfo 星河-获取微信用户的信息
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
func AlitripMerchantGalaxyWechatInfo(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatInfoAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
