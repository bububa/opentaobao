package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolRemoveitemAPIResponse 移除商品池里面的商品 API返回值
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabaHmMarketingItempoolRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolRemoveitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolRemoveitemAPIResponseModel is 移除商品池里面的商品 成功返回结果
type AlibabaHmMarketingItempoolRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolRemoveitemAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolRemoveitemAPIResponse
func GetAlibabaHmMarketingItempoolRemoveitemAPIResponse() *AlibabaHmMarketingItempoolRemoveitemAPIResponse {
	return poolAlibabaHmMarketingItempoolRemoveitemAPIResponse.Get().(*AlibabaHmMarketingItempoolRemoveitemAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolRemoveitemAPIResponse 将 AlibabaHmMarketingItempoolRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolRemoveitemAPIResponse(v *AlibabaHmMarketingItempoolRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolRemoveitemAPIResponse.Put(v)
}
