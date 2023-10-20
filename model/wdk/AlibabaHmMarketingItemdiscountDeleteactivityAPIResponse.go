package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse 删除商品特价活动 API返回值
// alibaba.hm.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItemdiscountDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItemdiscountDeleteactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingItemdiscountDeleteactivityAPIResponseModel is 删除商品特价活动 成功返回结果
type AlibabaHmMarketingItemdiscountDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse
func GetAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse() *AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse {
	return poolAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse.Get().(*AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse 将 AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse(v *AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountDeleteactivityAPIResponse.Put(v)
}
