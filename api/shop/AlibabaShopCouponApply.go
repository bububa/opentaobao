package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaShopCouponApply 通用店铺券领券接口
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
func AlibabaShopCouponApply(clt *core.SDKClient, req *shop.AlibabaShopCouponApplyAPIRequest, session string) (*shop.AlibabaShopCouponApplyAPIResponse, error) {
	var resp shop.AlibabaShopCouponApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
