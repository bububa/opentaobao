package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyReceiveCouponByActivity 按活动Id领取优惠券
// alitrip.merchant.galaxy.receive.coupon.by.activity
//
// 雅高小程序按活动Id领取优惠券
func AlitripMerchantGalaxyReceiveCouponByActivity(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
