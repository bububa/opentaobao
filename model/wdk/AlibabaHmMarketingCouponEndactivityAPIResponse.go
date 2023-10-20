package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponEndactivityAPIResponse 结束优惠券活动 API返回值
// alibaba.hm.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
type AlibabaHmMarketingCouponEndactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingCouponEndactivityAPIResponseModel
}

// AlibabaHmMarketingCouponEndactivityAPIResponseModel is 结束优惠券活动 成功返回结果
type AlibabaHmMarketingCouponEndactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_coupon_endactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
