package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingCouponQueryitemsAPIRequest
查询优惠券活动下的商品 API请求
alibaba.wdk.marketing.coupon.queryitems

查询优惠券活动下面的商品 */
type AlibabaWdkMarketingCouponQueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// New
