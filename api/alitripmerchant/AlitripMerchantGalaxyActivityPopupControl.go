package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivitypopupcontrol 营销弹屏疲劳度控制
// alitrip.merchant.galaxy.activity.popup.control
//
// 星河=营销弹屏疲劳控制
func Alitripmerchantgalaxyactivitypopupcontrol(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivitypopupcontrolAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivitypopupcontrolAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivitypopupcontrolAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
