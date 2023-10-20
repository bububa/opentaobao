package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountQueryactivityAPIResponse 查找特价活动 API返回值
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabaHmMarketingItemdiscountQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItemdiscountQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItemdiscountQueryactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingItemdiscountQueryactivityAPIResponseModel is 查找特价活动 成功返回结果
type AlibabaHmMarketingItemdiscountQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询特价活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItemdiscountQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItemdiscountQueryactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingItemdiscountQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountQueryactivityAPIResponse
func GetAlibabaHmMarketingItemdiscountQueryactivityAPIResponse() *AlibabaHmMarketingItemdiscountQueryactivityAPIResponse {
	return poolAlibabaHmMarketingItemdiscountQueryactivityAPIResponse.Get().(*AlibabaHmMarketingItemdiscountQueryactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingItemdiscountQueryactivityAPIResponse 将 AlibabaHmMarketingItemdiscountQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountQueryactivityAPIResponse(v *AlibabaHmMarketingItemdiscountQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountQueryactivityAPIResponse.Put(v)
}
