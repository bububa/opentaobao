package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponSpreadApply 普通发券
// alibaba.wdk.coupon.spread.apply
//
// 优惠券发放
func AlibabaWdkCouponSpreadApply(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaWdkCouponSpreadApplyAPIRequest, resp *promotion.AlibabaWdkCouponSpreadApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
