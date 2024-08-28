package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingCouponEndactivity 结束优惠券活动
// alibaba.hm.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
func AlibabaHmMarketingCouponEndactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingCouponEndactivityAPIRequest, resp *wdk.AlibabaHmMarketingCouponEndactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
