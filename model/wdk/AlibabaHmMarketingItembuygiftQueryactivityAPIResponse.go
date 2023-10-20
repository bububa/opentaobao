package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftQueryactivityAPIResponse 查询买赠活动 API返回值
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabaHmMarketingItembuygiftQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel is 查询买赠活动 成功返回结果
type AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItembuygiftQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItembuygiftQueryactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingItembuygiftQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftQueryactivityAPIResponse
func GetAlibabaHmMarketingItembuygiftQueryactivityAPIResponse() *AlibabaHmMarketingItembuygiftQueryactivityAPIResponse {
	return poolAlibabaHmMarketingItembuygiftQueryactivityAPIResponse.Get().(*AlibabaHmMarketingItembuygiftQueryactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingItembuygiftQueryactivityAPIResponse 将 AlibabaHmMarketingItembuygiftQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftQueryactivityAPIResponse(v *AlibabaHmMarketingItembuygiftQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftQueryactivityAPIResponse.Put(v)
}
