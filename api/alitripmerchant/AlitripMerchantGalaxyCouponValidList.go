package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxycouponvalidlist 用户有效优惠券列表
// alitrip.merchant.galaxy.coupon.valid.list
//
// 雅高小程序用户有效优惠券列表
func Alitripmerchantgalaxycouponvalidlist(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxycouponvalidlistAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxycouponvalidlistAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxycouponvalidlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
