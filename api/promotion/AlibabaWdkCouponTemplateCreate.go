package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponTemplateCreate 优惠券模版创建
// alibaba.wdk.coupon.template.create
//
// 开放给外部商家创建优惠券模版
func AlibabaWdkCouponTemplateCreate(clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateCreateAPIRequest, resp *promotion.AlibabaWdkCouponTemplateCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
