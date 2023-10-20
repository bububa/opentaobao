package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCouponInvalidList 用户失效优惠券列表
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
func AlitripMerchantGalaxyCouponInvalidList(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCouponInvalidListAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyCouponInvalidListAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyCouponInvalidListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
