package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftRemoveitemAPIResponse 移除买赠活动下的商品。【注意，此接口暂不支持并发！】 API返回值
// alibaba.hm.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
type AlibabaHmMarketingItembuygiftRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItembuygiftRemoveitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItembuygiftRemoveitemAPIResponseModel is 移除买赠活动下的商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaHmMarketingItembuygiftRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItembuygiftRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItembuygiftRemoveitemAPIResponse)
	},
}

// GetAlibabaHmMarketingItembuygiftRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftRemoveitemAPIResponse
func GetAlibabaHmMarketingItembuygiftRemoveitemAPIResponse() *AlibabaHmMarketingItembuygiftRemoveitemAPIResponse {
	return poolAlibabaHmMarketingItembuygiftRemoveitemAPIResponse.Get().(*AlibabaHmMarketingItembuygiftRemoveitemAPIResponse)
}

// ReleaseAlibabaHmMarketingItembuygiftRemoveitemAPIResponse 将 AlibabaHmMarketingItembuygiftRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftRemoveitemAPIResponse(v *AlibabaHmMarketingItembuygiftRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftRemoveitemAPIResponse.Put(v)
}
