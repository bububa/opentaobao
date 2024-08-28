package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaShopCouponApply 通用店铺券领券接口
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
func AlibabaShopCouponApply(ctx context.Context, clt *core.SDKClient, req *shop.AlibabaShopCouponApplyAPIRequest, resp *shop.AlibabaShopCouponApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
