package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivitymarketingpopup 营销弹屏
// alitrip.merchant.galaxy.activity.marketing.popup
//
// 星河=活动营销弹屏
func Alitripmerchantgalaxyactivitymarketingpopup(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivitymarketingpopupAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivitymarketingpopupAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivitymarketingpopupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
