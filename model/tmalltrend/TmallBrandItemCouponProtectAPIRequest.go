package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallBrandItemCouponProtectAPIRequest
全域新品店铺优惠券免除 API请求
tmall.brand.item.coupon.protect

全域新品店铺优惠券免除申请打标接口 */
type TmallBrandItemCouponProtectAPIRequest struct {
	model.Params
	// 天猫商品id
	_itemId int64
	// 店铺优惠券新品保护期档次:PERIOD_0D("0天"),     PERIOD_7D("7天"),     PERIOD_14D("14天"),     PERIOD_21D("21天")
	_protectionPeriod string
	// 天猫品牌id
	_brandId int64
}

// New
