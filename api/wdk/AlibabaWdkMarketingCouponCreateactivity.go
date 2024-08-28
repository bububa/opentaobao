package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingCouponCreateactivity 优惠券活动创建
// alibaba.wdk.marketing.coupon.createactivity
//
// 添加优惠券活动
func AlibabaWdkMarketingCouponCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingCouponCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
