package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingCouponEndactivity 结束优惠券活动
// alibaba.wdk.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
func AlibabaWdkMarketingCouponEndactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponEndactivityAPIRequest, resp *wdk.AlibabaWdkMarketingCouponEndactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
