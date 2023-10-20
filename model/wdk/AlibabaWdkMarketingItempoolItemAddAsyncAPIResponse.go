package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse 商品池新增商品 API返回值
// alibaba.wdk.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
type AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolItemAddAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolItemAddAsyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolItemAddAsyncAPIResponseModel is 商品池新增商品 成功返回结果
type AlibabaWdkMarketingItempoolItemAddAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolItemAddAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse
func GetAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse() *AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse {
	return poolAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse.Get().(*AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse 将 AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse(v *AlibabaWdkMarketingItempoolItemAddAsyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolItemAddAsyncAPIResponse.Put(v)
}
