package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkMarketingCouponSendma 发放匿名码
// alibaba.wdk.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
func AlibabaWdkMarketingCouponSendma(clt *core.SDKClient, req *promotion.AlibabaWdkMarketingCouponSendmaAPIRequest, resp *promotion.AlibabaWdkMarketingCouponSendmaAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
