package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeQueryitemAPIResponse 全场活动查询换购品 API返回值
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabaHmMarketingFullrangeQueryitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeQueryitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeQueryitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingFullrangeQueryitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingFullrangeQueryitemAPIResponseModel is 全场活动查询换购品 成功返回结果
type AlibabaHmMarketingFullrangeQueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeQueryitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingFullrangeQueryitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingFullrangeQueryitemAPIResponse)
	},
}

// GetAlibabaHmMarketingFullrangeQueryitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingFullrangeQueryitemAPIResponse
func GetAlibabaHmMarketingFullrangeQueryitemAPIResponse() *AlibabaHmMarketingFullrangeQueryitemAPIResponse {
	return poolAlibabaHmMarketingFullrangeQueryitemAPIResponse.Get().(*AlibabaHmMarketingFullrangeQueryitemAPIResponse)
}

// ReleaseAlibabaHmMarketingFullrangeQueryitemAPIResponse 将 AlibabaHmMarketingFullrangeQueryitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeQueryitemAPIResponse(v *AlibabaHmMarketingFullrangeQueryitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeQueryitemAPIResponse.Put(v)
}
