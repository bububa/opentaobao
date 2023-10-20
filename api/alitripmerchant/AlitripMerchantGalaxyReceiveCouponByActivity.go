package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyreceivecouponbyactivity 按活动Id领取优惠券
// alitrip.merchant.galaxy.receive.coupon.by.activity
//
// 雅高小程序按活动Id领取优惠券
func Alitripmerchantgalaxyreceivecouponbyactivity(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyreceivecouponbyactivityAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyreceivecouponbyactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
