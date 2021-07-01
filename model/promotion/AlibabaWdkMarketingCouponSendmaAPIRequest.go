package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingCouponSendmaAPIRequest
发放匿名码 API请求
alibaba.wdk.marketing.coupon.sendma

根据优惠券活动id打印单个匿名码 */
type AlibabaWdkMarketingCouponSendmaAPIRequest struct {
	model.Params
	// 发放匿名码入参
	_param0 *CommonActivityParam
}

// New
