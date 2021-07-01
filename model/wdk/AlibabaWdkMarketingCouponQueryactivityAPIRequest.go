package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingCouponQueryactivityAPIRequest
查询优惠券活动 API请求
alibaba.wdk.marketing.coupon.queryactivity

查询优惠券活动 */
type AlibabaWdkMarketingCouponQueryactivityAPIRequest struct {
	model.Params
	// 查询优惠券活动入参
	_param0 *CommonActivityParam
}

// New
