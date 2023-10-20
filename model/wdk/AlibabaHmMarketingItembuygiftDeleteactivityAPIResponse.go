package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse 删除买赠活动 API返回值
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItembuygiftDeleteactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingItembuygiftDeleteactivityAPIResponseModel is 删除买赠活动 成功返回结果
type AlibabaHmMarketingItembuygiftDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItembuygiftDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse
func GetAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse() *AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse {
	return poolAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse.Get().(*AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse 将 AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse(v *AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItembuygiftDeleteactivityAPIResponse.Put(v)
}
