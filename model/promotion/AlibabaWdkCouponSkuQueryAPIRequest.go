package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponSkuQueryAPIRequest
优惠券商品查询 API请求
alibaba.wdk.coupon.sku.query

优惠券商品查询 */
type AlibabaWdkCouponSkuQueryAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateItemQueryRequest *CouponTemplateItemQueryRequest
}

// New
