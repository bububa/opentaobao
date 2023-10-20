package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolItemAddAsyncAPIResponse 商品池新增商品 API返回值
// alibaba.hm.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
type AlibabaHmMarketingItempoolItemAddAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolItemAddAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolItemAddAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolItemAddAsyncAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolItemAddAsyncAPIResponseModel is 商品池新增商品 成功返回结果
type AlibabaHmMarketingItempoolItemAddAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolItemAddAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolItemAddAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolItemAddAsyncAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolItemAddAsyncAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolItemAddAsyncAPIResponse
func GetAlibabaHmMarketingItempoolItemAddAsyncAPIResponse() *AlibabaHmMarketingItempoolItemAddAsyncAPIResponse {
	return poolAlibabaHmMarketingItempoolItemAddAsyncAPIResponse.Get().(*AlibabaHmMarketingItempoolItemAddAsyncAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolItemAddAsyncAPIResponse 将 AlibabaHmMarketingItempoolItemAddAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolItemAddAsyncAPIResponse(v *AlibabaHmMarketingItempoolItemAddAsyncAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolItemAddAsyncAPIResponse.Put(v)
}
