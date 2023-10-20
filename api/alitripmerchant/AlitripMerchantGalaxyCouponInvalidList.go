package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxycouponinvalidlist 用户失效优惠券列表
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
func Alitripmerchantgalaxycouponinvalidlist(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxycouponinvalidlistAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxycouponinvalidlistAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxycouponinvalidlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
