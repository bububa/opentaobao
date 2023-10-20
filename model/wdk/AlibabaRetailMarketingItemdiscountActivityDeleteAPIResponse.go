package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse 删除单品特价活动【同城零售】 API返回值
// alibaba.retail.marketing.itemdiscount.activity.delete
//
// 同城零售单品特价活动删除
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponseModel is 删除单品特价活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse)
	},
}

// GetAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse
func GetAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse() *AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse {
	return poolAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse.Get().(*AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse 将 AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse(v *AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse.Put(v)
}
