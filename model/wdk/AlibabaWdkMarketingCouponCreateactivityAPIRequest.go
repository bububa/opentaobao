package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingCouponCreateactivityAPIRequest
优惠券活动创建 API请求
alibaba.wdk.marketing.coupon.createactivity

添加优惠券活动 */
type AlibabaWdkMarketingCouponCreateactivityAPIRequest struct {
	model.Params
	// 创建优惠券活动请求入参
	_param *CouponActivity
}

// New
