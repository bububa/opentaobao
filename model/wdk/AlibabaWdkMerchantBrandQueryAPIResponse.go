package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantBrandQueryAPIResponse 品牌查询接口 API返回值
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
type AlibabaWdkMerchantBrandQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantBrandQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantBrandQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMerchantBrandQueryAPIResponseModel).Reset()
}

// AlibabaWdkMerchantBrandQueryAPIResponseModel is 品牌查询接口 成功返回结果
type AlibabaWdkMerchantBrandQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_brand_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkMerchantBrandQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantBrandQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMerchantBrandQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantBrandQueryAPIResponse)
	},
}

// GetAlibabaWdkMerchantBrandQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkMerchantBrandQueryAPIResponse
func GetAlibabaWdkMerchantBrandQueryAPIResponse() *AlibabaWdkMerchantBrandQueryAPIResponse {
	return poolAlibabaWdkMerchantBrandQueryAPIResponse.Get().(*AlibabaWdkMerchantBrandQueryAPIResponse)
}

// ReleaseAlibabaWdkMerchantBrandQueryAPIResponse 将 AlibabaWdkMerchantBrandQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMerchantBrandQueryAPIResponse(v *AlibabaWdkMerchantBrandQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMerchantBrandQueryAPIResponse.Put(v)
}
