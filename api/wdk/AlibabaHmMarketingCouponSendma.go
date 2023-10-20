package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingCouponSendma 发放匿名码
// alibaba.hm.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
func AlibabaHmMarketingCouponSendma(clt *core.SDKClient, req *wdk.AlibabaHmMarketingCouponSendmaAPIRequest, resp *wdk.AlibabaHmMarketingCouponSendmaAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
