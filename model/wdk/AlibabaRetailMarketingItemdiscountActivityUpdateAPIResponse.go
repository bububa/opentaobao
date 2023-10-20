package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse 更新单品特价活动【同城零售】 API返回值
// alibaba.retail.marketing.itemdiscount.activity.update
//
// 同城零售单品特价活动更新
type AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponseModel is 更新单品特价活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse)
	},
}

// GetAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse
func GetAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse() *AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse {
	return poolAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse.Get().(*AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse 将 AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse(v *AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse.Put(v)
}
