package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponAbandon 废券
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
func AlibabaWdkCouponAbandon(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaWdkCouponAbandonAPIRequest, resp *promotion.AlibabaWdkCouponAbandonAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
