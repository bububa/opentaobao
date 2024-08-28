package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponTemplateUpdate 优惠券模版修改
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
func AlibabaWdkCouponTemplateUpdate(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateUpdateAPIRequest, resp *promotion.AlibabaWdkCouponTemplateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
