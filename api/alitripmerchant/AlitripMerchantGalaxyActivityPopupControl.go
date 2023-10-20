package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityPopupControl 营销弹屏疲劳度控制
// alitrip.merchant.galaxy.activity.popup.control
//
// 星河=营销弹屏疲劳控制
func AlitripMerchantGalaxyActivityPopupControl(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityPopupControlAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityPopupControlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
