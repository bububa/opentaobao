package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponSkuAddAPIRequest
优惠券商品增加 API请求
alibaba.wdk.coupon.sku.add

优惠券商品增加 */
type AlibabaWdkCouponSkuAddAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// New
