package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftAdditemAPIResponse 增加买赠活动商品。【注意，此接口暂不支持并发！】 API返回值
// alibaba.hm.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
type AlibabaHmMarketingItembuygiftAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItembuygiftAdditemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItembuygiftAdditemAPIResponseModel is 增加买赠活动商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaHmMarketingItembuygiftAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItembuygiftAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItembuygiftAdditemAPIResponse)
	},
}

// GetAlibabaHmMarketingItembuygiftAdditemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftAdditemAPIResponse
func GetAlibabaHmMarketingItembuygiftAdditemAPIResponse() *AlibabaHmMarketingItembuygiftAdditemAPIResponse {
	return poolAlibabaHmMarketingItembuygiftAdditemAPIResponse.Get().(*AlibabaHmMarketingItembuygiftAdditemAPIResponse)
}

// ReleaseAlibabaHmMarketingItembuygiftAdditemAPIResponse 将 AlibabaHmMarketingItembuygiftAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftAdditemAPIResponse(v *AlibabaHmMarketingItembuygiftAdditemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftAdditemAPIResponse.Put(v)
}
