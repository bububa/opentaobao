package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponSendmaAPIResponse 发放匿名码 API返回值
// alibaba.hm.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
type AlibabaHmMarketingCouponSendmaAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingCouponSendmaAPIResponseModel
}

// AlibabaHmMarketingCouponSendmaAPIResponseModel is 发放匿名码 成功返回结果
type AlibabaHmMarketingCouponSendmaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_coupon_sendma_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发放匿名码返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
