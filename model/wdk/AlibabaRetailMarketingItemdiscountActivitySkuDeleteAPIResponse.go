package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse 删除特价活动商品 API返回值
// alibaba.retail.marketing.itemdiscount.activity.sku.delete
//
// 删除活动商品信息【同城零售】
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponseModel is 删除特价活动商品 成功返回结果
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse)
	},
}

// GetAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse
func GetAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse() *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse {
	return poolAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse.Get().(*AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse 将 AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse(v *AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse.Put(v)
}
