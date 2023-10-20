package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponTemplateQueryAPIResponse 优惠券模版查询 API返回值
// alibaba.wdk.coupon.template.query
//
// 优惠券模版查询
type AlibabaWdkCouponTemplateQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponTemplateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponTemplateQueryAPIResponseModel).Reset()
}

// AlibabaWdkCouponTemplateQueryAPIResponseModel is 优惠券模版查询 成功返回结果
type AlibabaWdkCouponTemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaWdkCouponTemplateQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponTemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponTemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponTemplateQueryAPIResponse)
	},
}

// GetAlibabaWdkCouponTemplateQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponTemplateQueryAPIResponse
func GetAlibabaWdkCouponTemplateQueryAPIResponse() *AlibabaWdkCouponTemplateQueryAPIResponse {
	return poolAlibabaWdkCouponTemplateQueryAPIResponse.Get().(*AlibabaWdkCouponTemplateQueryAPIResponse)
}

// ReleaseAlibabaWdkCouponTemplateQueryAPIResponse 将 AlibabaWdkCouponTemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponTemplateQueryAPIResponse(v *AlibabaWdkCouponTemplateQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponTemplateQueryAPIResponse.Put(v)
}
