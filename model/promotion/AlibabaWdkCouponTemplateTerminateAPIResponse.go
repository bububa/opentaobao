package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateTerminateAPIResponse 优惠券模版终止 API返回值
// alibaba.wdk.coupon.template.terminate
//
// 优惠券模版终止
type AlibabaWdkCouponTemplateTerminateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponTemplateTerminateAPIResponseModel
}

// AlibabaWdkCouponTemplateTerminateAPIResponseModel is 优惠券模版终止 成功返回结果
type AlibabaWdkCouponTemplateTerminateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_template_terminate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponTemplateTerminateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
