package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatInfo 星河-获取微信用户的信息
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
func AlitripMerchantGalaxyWechatInfo(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatInfoAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
