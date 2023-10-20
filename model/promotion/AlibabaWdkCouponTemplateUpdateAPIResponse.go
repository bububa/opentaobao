package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateUpdateAPIResponse 优惠券模版修改 API返回值
// alibaba.wdk.coupon.template.update
//
// 优惠券模版修改
type AlibabaWdkCouponTemplateUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponTemplateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponTemplateUpdateAPIResponseModel).Reset()
}

// AlibabaWdkCouponTemplateUpdateAPIResponseModel is 优惠券模版修改 成功返回结果
type AlibabaWdkCouponTemplateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_template_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponTemplateUpdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponTemplateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateUpdateAPIResponse)
	},
}

// GetAlibabaWdkCouponTemplateUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponTemplateUpdateAPIResponse
func GetAlibabaWdkCouponTemplateUpdateAPIResponse() *AlibabaWdkCouponTemplateUpdateAPIResponse {
	return poolAlibabaWdkCouponTemplateUpdateAPIResponse.Get().(*AlibabaWdkCouponTemplateUpdateAPIResponse)
}

// ReleaseAlibabaWdkCouponTemplateUpdateAPIResponse 将 AlibabaWdkCouponTemplateUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponTemplateUpdateAPIResponse(v *AlibabaWdkCouponTemplateUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponTemplateUpdateAPIResponse.Put(v)
}
