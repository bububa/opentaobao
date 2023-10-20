package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponTemplateTerminate 优惠券模版终止
// alibaba.wdk.coupon.template.terminate
//
// 优惠券模版终止
func AlibabaWdkCouponTemplateTerminate(clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateTerminateAPIRequest, resp *promotion.AlibabaWdkCouponTemplateTerminateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
