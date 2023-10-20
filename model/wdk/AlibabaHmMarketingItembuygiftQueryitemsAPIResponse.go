package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftQueryitemsAPIResponse 查询买赠活动下的商品 API返回值
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabaHmMarketingItembuygiftQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItembuygiftQueryitemsAPIResponseModel).Reset()
}

// AlibabaHmMarketingItembuygiftQueryitemsAPIResponseModel is 查询买赠活动下的商品 成功返回结果
type AlibabaHmMarketingItembuygiftQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItembuygiftQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItembuygiftQueryitemsAPIResponse)
	},
}

// GetAlibabaHmMarketingItembuygiftQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftQueryitemsAPIResponse
func GetAlibabaHmMarketingItembuygiftQueryitemsAPIResponse() *AlibabaHmMarketingItembuygiftQueryitemsAPIResponse {
	return poolAlibabaHmMarketingItembuygiftQueryitemsAPIResponse.Get().(*AlibabaHmMarketingItembuygiftQueryitemsAPIResponse)
}

// ReleaseAlibabaHmMarketingItembuygiftQueryitemsAPIResponse 将 AlibabaHmMarketingItembuygiftQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftQueryitemsAPIResponse(v *AlibabaHmMarketingItembuygiftQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftQueryitemsAPIResponse.Put(v)
}
