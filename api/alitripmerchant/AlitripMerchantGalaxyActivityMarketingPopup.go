package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityMarketingPopup 营销弹屏
// alitrip.merchant.galaxy.activity.marketing.popup
//
// 星河=活动营销弹屏
func AlitripMerchantGalaxyActivityMarketingPopup(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityMarketingPopupAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityMarketingPopupAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
