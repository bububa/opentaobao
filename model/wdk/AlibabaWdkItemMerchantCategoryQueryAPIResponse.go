package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantCategoryQueryAPIResponse 查询商品的商家叶子类目 API返回值
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
type AlibabaWdkItemMerchantCategoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantCategoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantCategoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemMerchantCategoryQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemMerchantCategoryQueryAPIResponseModel is 查询商品的商家叶子类目 成功返回结果
type AlibabaWdkItemMerchantCategoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchant_category_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *WdkOpenSkuMerchantCatServiceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantCategoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemMerchantCategoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantCategoryQueryAPIResponse)
	},
}

// GetAlibabaWdkItemMerchantCategoryQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemMerchantCategoryQueryAPIResponse
func GetAlibabaWdkItemMerchantCategoryQueryAPIResponse() *AlibabaWdkItemMerchantCategoryQueryAPIResponse {
	return poolAlibabaWdkItemMerchantCategoryQueryAPIResponse.Get().(*AlibabaWdkItemMerchantCategoryQueryAPIResponse)
}

// ReleaseAlibabaWdkItemMerchantCategoryQueryAPIResponse 将 AlibabaWdkItemMerchantCategoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemMerchantCategoryQueryAPIResponse(v *AlibabaWdkItemMerchantCategoryQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemMerchantCategoryQueryAPIResponse.Put(v)
}
