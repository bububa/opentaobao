package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* AlibabaWdkCouponTemplateTerminate
优惠券模版终止
alibaba.wdk.coupon.template.terminate

优惠券模版终止 */
func AlibabaWdkCouponTemplateTerminate(clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateTerminateAPIRequest, session string) (*promotion.AlibabaWdkCouponTemplateTerminateAPIResponse, error) {
	var resp promotion.AlibabaWdkCouponTemplateTerminateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
