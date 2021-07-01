package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponSkuRemoveAPIRequest
优惠券商品删除 API请求
alibaba.wdk.coupon.sku.remove

优惠券商品删除 */
type AlibabaWdkCouponSkuRemoveAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemRequest *CouponTemplateItemRequest
}

// New
