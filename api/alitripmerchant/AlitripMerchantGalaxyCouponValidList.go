package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCouponValidList 用户有效优惠券列表
// alitrip.merchant.galaxy.coupon.valid.list
//
// 雅高小程序用户有效优惠券列表
func AlitripMerchantGalaxyCouponValidList(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCouponValidListAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyCouponValidListAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyCouponValidListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
