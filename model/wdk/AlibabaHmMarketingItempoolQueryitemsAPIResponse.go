package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolQueryitemsAPIResponse 查询商品池活动下面的商品 API返回值
// alibaba.hm.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabaHmMarketingItempoolQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolQueryitemsAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolQueryitemsAPIResponseModel is 查询商品池活动下面的商品 成功返回结果
type AlibabaHmMarketingItempoolQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolQueryitemsAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolQueryitemsAPIResponse
func GetAlibabaHmMarketingItempoolQueryitemsAPIResponse() *AlibabaHmMarketingItempoolQueryitemsAPIResponse {
	return poolAlibabaHmMarketingItempoolQueryitemsAPIResponse.Get().(*AlibabaHmMarketingItempoolQueryitemsAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolQueryitemsAPIResponse 将 AlibabaHmMarketingItempoolQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolQueryitemsAPIResponse(v *AlibabaHmMarketingItempoolQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolQueryitemsAPIResponse.Put(v)
}
