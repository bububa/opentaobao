package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponTemplateTerminateAPIRequest
优惠券模版终止 API请求
alibaba.wdk.coupon.template.terminate

优惠券模版终止 */
type AlibabaWdkCouponTemplateTerminateAPIRequest struct {
	model.Params
	// 参数
	_paramCouponTemplateTerminateRequest *CouponTemplateTerminateRequest
}

// New
