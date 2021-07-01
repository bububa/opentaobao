package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingCouponEndactivityAPIRequest
结束优惠券活动 API请求
alibaba.wdk.marketing.coupon.endactivity

结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用 */
type AlibabaWdkMarketingCouponEndactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// New
