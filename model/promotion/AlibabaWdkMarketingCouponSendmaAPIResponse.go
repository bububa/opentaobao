package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponSendmaAPIResponse 发放匿名码 API返回值
// alibaba.wdk.marketing.coupon.sendma
//
// 根据优惠券活动id打印单个匿名码
type AlibabaWdkMarketingCouponSendmaAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingCouponSendmaAPIResponseModel
}

// AlibabaWdkMarketingCouponSendmaAPIResponseModel is 发放匿名码 成功返回结果
type AlibabaWdkMarketingCouponSendmaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_sendma_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发放匿名码返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
