package tmalltrend

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmalltrend"
)

// TmallBrandItemCouponProtect 全域新品店铺优惠券免除
// tmall.brand.item.coupon.protect
//
// 全域新品店铺优惠券免除申请打标接口
func TmallBrandItemCouponProtect(clt *core.SDKClient, req *tmalltrend.TmallBrandItemCouponProtectAPIRequest, resp *tmalltrend.TmallBrandItemCouponProtectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
