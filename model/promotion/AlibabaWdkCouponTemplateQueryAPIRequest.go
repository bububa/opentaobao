package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponTemplateQueryAPIRequest
优惠券模版查询 API请求
alibaba.wdk.coupon.template.query

优惠券模版查询 */
type AlibabaWdkCouponTemplateQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramCouponTemplateQueryRequest *CouponTemplateQueryRequest
}

// New
