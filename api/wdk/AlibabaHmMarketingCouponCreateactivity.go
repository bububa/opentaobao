package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingCouponCreateactivity 优惠券活动创建
// alibaba.hm.marketing.coupon.createactivity
//
// 添加优惠券活动
func AlibabaHmMarketingCouponCreateactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingCouponCreateactivityAPIRequest, resp *wdk.AlibabaHmMarketingCouponCreateactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
