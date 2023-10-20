package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyordercouponvalidate 携带券的试单接口
// alitrip.merchant.galaxy.order.coupon.validate
//
// 试单时可以使用优惠券，返回一个原价，一个折扣价
func Alitripmerchantgalaxyordercouponvalidate(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyordercouponvalidateAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyordercouponvalidateAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyordercouponvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
