package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateCreateAPIResponse 优惠券模版创建 API返回值
// alibaba.wdk.coupon.template.create
//
// 开放给外部商家创建优惠券模版
type AlibabaWdkCouponTemplateCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponTemplateCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponTemplateCreateAPIResponseModel).Reset()
}

// AlibabaWdkCouponTemplateCreateAPIResponseModel is 优惠券模版创建 成功返回结果
type AlibabaWdkCouponTemplateCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_template_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponTemplateCreateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponTemplateCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateCreateAPIResponse)
	},
}

// GetAlibabaWdkCouponTemplateCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponTemplateCreateAPIResponse
func GetAlibabaWdkCouponTemplateCreateAPIResponse() *AlibabaWdkCouponTemplateCreateAPIResponse {
	return poolAlibabaWdkCouponTemplateCreateAPIResponse.Get().(*AlibabaWdkCouponTemplateCreateAPIResponse)
}

// ReleaseAlibabaWdkCouponTemplateCreateAPIResponse 将 AlibabaWdkCouponTemplateCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponTemplateCreateAPIResponse(v *AlibabaWdkCouponTemplateCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponTemplateCreateAPIResponse.Put(v)
}
