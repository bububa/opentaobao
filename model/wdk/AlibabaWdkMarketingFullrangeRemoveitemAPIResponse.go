package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeRemoveitemAPIResponse 全场活动删除购品 API返回值
// alibaba.wdk.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabaWdkMarketingFullrangeRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel is 全场活动删除购品 成功返回结果
type AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingFullrangeRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingFullrangeRemoveitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingFullrangeRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeRemoveitemAPIResponse
func GetAlibabaWdkMarketingFullrangeRemoveitemAPIResponse() *AlibabaWdkMarketingFullrangeRemoveitemAPIResponse {
	return poolAlibabaWdkMarketingFullrangeRemoveitemAPIResponse.Get().(*AlibabaWdkMarketingFullrangeRemoveitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingFullrangeRemoveitemAPIResponse 将 AlibabaWdkMarketingFullrangeRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeRemoveitemAPIResponse(v *AlibabaWdkMarketingFullrangeRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeRemoveitemAPIResponse.Put(v)
}
