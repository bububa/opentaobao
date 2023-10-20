package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingCouponSendma 发放匿名码
// alibaba.hm.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
func AlibabaHmMarketingCouponSendma(clt *core.SDKClient, req *wdk.AlibabaHmMarketingCouponSendmaAPIRequest, session string) (*wdk.AlibabaHmMarketingCouponSendmaAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingCouponSendmaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
