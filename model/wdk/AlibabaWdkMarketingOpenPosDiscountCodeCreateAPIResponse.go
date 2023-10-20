package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse pos一物一码创建 API返回值
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
type AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponseModel).Reset()
}

// AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponseModel is pos一物一码创建 成功返回结果
type AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_pos_discount_code_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse)
	},
}

// GetAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse
func GetAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse() *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse {
	return poolAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse.Get().(*AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse)
}

// ReleaseAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse 将 AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse(v *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIResponse.Put(v)
}
