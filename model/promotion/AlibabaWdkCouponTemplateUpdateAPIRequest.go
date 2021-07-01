package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponTemplateUpdateAPIRequest
优惠券模版修改 API请求
alibaba.wdk.coupon.template.update

优惠券模版修改 */
type AlibabaWdkCouponTemplateUpdateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// New
