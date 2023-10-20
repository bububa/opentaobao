package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponTemplateQuery 优惠券模版查询
// alibaba.wdk.coupon.template.query
//
// 优惠券模版查询
func AlibabaWdkCouponTemplateQuery(clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateQueryAPIRequest, resp *promotion.AlibabaWdkCouponTemplateQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
