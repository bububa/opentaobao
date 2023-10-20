package promotion

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateTerminateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponTemplateTerminateAPIResponseModel).Reset()
}

// AlibabaWdkCouponTemplateTerminateAPIResponseModel is 优惠券模版终止 成功返回结果
type AlibabaWdkCouponTemplateTerminateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_template_terminate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponTemplateTerminateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateTerminateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponTemplateTerminateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateTerminateAPIResponse)
	},
}

// GetAlibabaWdkCouponTemplateTerminateAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponTemplateTerminateAPIResponse
func GetAlibabaWdkCouponTemplateTerminateAPIResponse() *AlibabaWdkCouponTemplateTerminateAPIResponse {
	return poolAlibabaWdkCouponTemplateTerminateAPIResponse.Get().(*AlibabaWdkCouponTemplateTerminateAPIResponse)
}

// ReleaseAlibabaWdkCouponTemplateTerminateAPIResponse 将 AlibabaWdkCouponTemplateTerminateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponTemplateTerminateAPIResponse(v *AlibabaWdkCouponTemplateTerminateAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponTemplateTerminateAPIResponse.Put(v)
}
