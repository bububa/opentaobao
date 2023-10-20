package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivitycouponlist 用户领券中心列表
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
func Alitripmerchantgalaxyactivitycouponlist(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivitycouponlistAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivitycouponlistAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivitycouponlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
