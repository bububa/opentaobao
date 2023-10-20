package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolRemoveitemAPIResponse 移除商品池里面的商品 API返回值
// alibaba.wdk.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabaWdkMarketingItempoolRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolRemoveitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolRemoveitemAPIResponseModel is 移除商品池里面的商品 成功返回结果
type AlibabaWdkMarketingItempoolRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolRemoveitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolRemoveitemAPIResponse
func GetAlibabaWdkMarketingItempoolRemoveitemAPIResponse() *AlibabaWdkMarketingItempoolRemoveitemAPIResponse {
	return poolAlibabaWdkMarketingItempoolRemoveitemAPIResponse.Get().(*AlibabaWdkMarketingItempoolRemoveitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolRemoveitemAPIResponse 将 AlibabaWdkMarketingItempoolRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolRemoveitemAPIResponse(v *AlibabaWdkMarketingItempoolRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolRemoveitemAPIResponse.Put(v)
}
