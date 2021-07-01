package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponTemplateCreateAPIRequest
优惠券模版创建 API请求
alibaba.wdk.coupon.template.create

开放给外部商家创建优惠券模版 */
type AlibabaWdkCouponTemplateCreateAPIRequest struct {
	model.Params
	// 请求
	_paramCouponTemplateOperateRequest *CouponTemplateOperateRequest
}

// New
