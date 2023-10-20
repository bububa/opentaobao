package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantItemQueryAPIResponse 商家商品查询 API返回值
// alibaba.wdk.merchant.item.query
//
// 商家商品查询
type AlibabaWdkMerchantItemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMerchantItemQueryAPIResponseModel).Reset()
}

// AlibabaWdkMerchantItemQueryAPIResponseModel is 商家商品查询 成功返回结果
type AlibabaWdkMerchantItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkMerchantItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMerchantItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantItemQueryAPIResponse)
	},
}

// GetAlibabaWdkMerchantItemQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkMerchantItemQueryAPIResponse
func GetAlibabaWdkMerchantItemQueryAPIResponse() *AlibabaWdkMerchantItemQueryAPIResponse {
	return poolAlibabaWdkMerchantItemQueryAPIResponse.Get().(*AlibabaWdkMerchantItemQueryAPIResponse)
}

// ReleaseAlibabaWdkMerchantItemQueryAPIResponse 将 AlibabaWdkMerchantItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMerchantItemQueryAPIResponse(v *AlibabaWdkMerchantItemQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMerchantItemQueryAPIResponse.Put(v)
}
