package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse 创建单品特价活动【同城零售】 API返回值
// alibaba.retail.marketing.itemdiscount.activity.create
//
// 同城零售单品特价活动创建
type AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItemdiscountActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItemdiscountActivityCreateAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItemdiscountActivityCreateAPIResponseModel is 创建单品特价活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItemdiscountActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse)
	},
}

// GetAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse
func GetAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse() *AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse {
	return poolAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse.Get().(*AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse 将 AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse(v *AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivityCreateAPIResponse.Put(v)
}
