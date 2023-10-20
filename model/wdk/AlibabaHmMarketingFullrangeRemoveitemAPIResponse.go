package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeRemoveitemAPIResponse 全场活动删除购品 API返回值
// alibaba.hm.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabaHmMarketingFullrangeRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel is 全场活动删除购品 成功返回结果
type AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingFullrangeRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingFullrangeRemoveitemAPIResponse)
	},
}

// GetAlibabaHmMarketingFullrangeRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingFullrangeRemoveitemAPIResponse
func GetAlibabaHmMarketingFullrangeRemoveitemAPIResponse() *AlibabaHmMarketingFullrangeRemoveitemAPIResponse {
	return poolAlibabaHmMarketingFullrangeRemoveitemAPIResponse.Get().(*AlibabaHmMarketingFullrangeRemoveitemAPIResponse)
}

// ReleaseAlibabaHmMarketingFullrangeRemoveitemAPIResponse 将 AlibabaHmMarketingFullrangeRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeRemoveitemAPIResponse(v *AlibabaHmMarketingFullrangeRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeRemoveitemAPIResponse.Put(v)
}
