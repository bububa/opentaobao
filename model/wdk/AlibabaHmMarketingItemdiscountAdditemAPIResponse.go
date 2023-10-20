package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountAdditemAPIResponse 报名特价商品 API返回值
// alibaba.hm.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
type AlibabaHmMarketingItemdiscountAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItemdiscountAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItemdiscountAdditemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItemdiscountAdditemAPIResponseModel is 报名特价商品 成功返回结果
type AlibabaHmMarketingItemdiscountAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItemdiscountAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItemdiscountAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItemdiscountAdditemAPIResponse)
	},
}

// GetAlibabaHmMarketingItemdiscountAdditemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItemdiscountAdditemAPIResponse
func GetAlibabaHmMarketingItemdiscountAdditemAPIResponse() *AlibabaHmMarketingItemdiscountAdditemAPIResponse {
	return poolAlibabaHmMarketingItemdiscountAdditemAPIResponse.Get().(*AlibabaHmMarketingItemdiscountAdditemAPIResponse)
}

// ReleaseAlibabaHmMarketingItemdiscountAdditemAPIResponse 将 AlibabaHmMarketingItemdiscountAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItemdiscountAdditemAPIResponse(v *AlibabaHmMarketingItemdiscountAdditemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItemdiscountAdditemAPIResponse.Put(v)
}
