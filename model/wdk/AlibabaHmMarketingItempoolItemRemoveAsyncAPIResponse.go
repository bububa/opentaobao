package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse 商品池删除商品 API返回值
// alibaba.hm.marketing.itempool.item.remove.async
//
// 新模型下删除商品
type AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel is 商品池删除商品 成功返回结果
type AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse
func GetAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse() *AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse {
	return poolAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse.Get().(*AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse 将 AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse(v *AlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolItemRemoveAsyncAPIResponse.Put(v)
}
