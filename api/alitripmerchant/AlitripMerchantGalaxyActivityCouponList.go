package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityCouponList 用户领券中心列表
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
func AlitripMerchantGalaxyActivityCouponList(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityCouponListAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyActivityCouponListAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyActivityCouponListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
