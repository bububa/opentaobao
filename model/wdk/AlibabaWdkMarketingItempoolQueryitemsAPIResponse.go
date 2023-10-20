package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolQueryitemsAPIResponse 查询商品池活动下的商品 API返回值
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabaWdkMarketingItempoolQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolQueryitemsAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolQueryitemsAPIResponseModel is 查询商品池活动下的商品 成功返回结果
type AlibabaWdkMarketingItempoolQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolQueryitemsAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolQueryitemsAPIResponse
func GetAlibabaWdkMarketingItempoolQueryitemsAPIResponse() *AlibabaWdkMarketingItempoolQueryitemsAPIResponse {
	return poolAlibabaWdkMarketingItempoolQueryitemsAPIResponse.Get().(*AlibabaWdkMarketingItempoolQueryitemsAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolQueryitemsAPIResponse 将 AlibabaWdkMarketingItempoolQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolQueryitemsAPIResponse(v *AlibabaWdkMarketingItempoolQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolQueryitemsAPIResponse.Put(v)
}
