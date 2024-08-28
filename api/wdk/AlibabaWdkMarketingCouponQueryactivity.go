package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingCouponQueryactivity 查询优惠券活动
// alibaba.wdk.marketing.coupon.queryactivity
//
// 查询优惠券活动
func AlibabaWdkMarketingCouponQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingCouponQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
