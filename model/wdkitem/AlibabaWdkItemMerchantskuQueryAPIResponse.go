package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantskuQueryAPIResponse 商家商品信息查询 API返回值
// alibaba.wdk.item.merchantsku.query
//
// 商家商品信息查询
type AlibabaWdkItemMerchantskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantskuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantskuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemMerchantskuQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemMerchantskuQueryAPIResponseModel is 商家商品信息查询 成功返回结果
type AlibabaWdkItemMerchantskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantskuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemMerchantskuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantskuQueryAPIResponse)
	},
}

// GetAlibabaWdkItemMerchantskuQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemMerchantskuQueryAPIResponse
func GetAlibabaWdkItemMerchantskuQueryAPIResponse() *AlibabaWdkItemMerchantskuQueryAPIResponse {
	return poolAlibabaWdkItemMerchantskuQueryAPIResponse.Get().(*AlibabaWdkItemMerchantskuQueryAPIResponse)
}

// ReleaseAlibabaWdkItemMerchantskuQueryAPIResponse 将 AlibabaWdkItemMerchantskuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemMerchantskuQueryAPIResponse(v *AlibabaWdkItemMerchantskuQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemMerchantskuQueryAPIResponse.Put(v)
}
