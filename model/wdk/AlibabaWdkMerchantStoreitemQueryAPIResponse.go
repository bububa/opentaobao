package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantStoreitemQueryAPIResponse 门店商品信心查询 API返回值
// alibaba.wdk.merchant.storeitem.query
//
// 门店商品信心查询
type AlibabaWdkMerchantStoreitemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantStoreitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantStoreitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMerchantStoreitemQueryAPIResponseModel).Reset()
}

// AlibabaWdkMerchantStoreitemQueryAPIResponseModel is 门店商品信心查询 成功返回结果
type AlibabaWdkMerchantStoreitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_storeitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkMerchantStoreitemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantStoreitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMerchantStoreitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantStoreitemQueryAPIResponse)
	},
}

// GetAlibabaWdkMerchantStoreitemQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkMerchantStoreitemQueryAPIResponse
func GetAlibabaWdkMerchantStoreitemQueryAPIResponse() *AlibabaWdkMerchantStoreitemQueryAPIResponse {
	return poolAlibabaWdkMerchantStoreitemQueryAPIResponse.Get().(*AlibabaWdkMerchantStoreitemQueryAPIResponse)
}

// ReleaseAlibabaWdkMerchantStoreitemQueryAPIResponse 将 AlibabaWdkMerchantStoreitemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMerchantStoreitemQueryAPIResponse(v *AlibabaWdkMerchantStoreitemQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMerchantStoreitemQueryAPIResponse.Put(v)
}
