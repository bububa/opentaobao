package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyShareGet 星河-获取小程序分享文案和图片
// alitrip.merchant.galaxy.share.get
//
// 获取 雅高微信小程序分享素材文案和图片。
func AlitripMerchantGalaxyShareGet(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyShareGetAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyShareGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
