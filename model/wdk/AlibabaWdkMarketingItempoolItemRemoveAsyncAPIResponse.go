package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse 商品池删除商品 API返回值
// alibaba.wdk.marketing.itempool.item.remove.async
//
// 新模型下删除商品
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponseModel is 商品池删除商品 成功返回结果
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse
func GetAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse() *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse {
	return poolAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse.Get().(*AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse 将 AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse(v *AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse.Put(v)
}
