package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountQueryitemsAPIResponse 查询特价商品 API返回值
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
type AlibabaHmMarketingItemdiscountQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItemdiscountQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItemdiscountQueryitemsAPIResponseModel).Reset()
}

// AlibabaHmMarketingItemdiscountQueryitemsAPIResponseModel is 查询特价商品 成功返回结果
type AlibabaHmMarketingItemdiscountQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItemdiscountQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItemdiscountQueryitemsAPIResponse)
	},
}

// GetAlibabaHmMarketingItemdiscountQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountQueryitemsAPIResponse
func GetAlibabaHmMarketingItemdiscountQueryitemsAPIResponse() *AlibabaHmMarketingItemdiscountQueryitemsAPIResponse {
	return poolAlibabaHmMarketingItemdiscountQueryitemsAPIResponse.Get().(*AlibabaHmMarketingItemdiscountQueryitemsAPIResponse)
}

// ReleaseAlibabaHmMarketingItemdiscountQueryitemsAPIResponse 将 AlibabaHmMarketingItemdiscountQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountQueryitemsAPIResponse(v *AlibabaHmMarketingItemdiscountQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountQueryitemsAPIResponse.Put(v)
}
