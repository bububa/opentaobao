package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolStairQueryitemAPIResponse 换购商品查询 API返回值
// alibaba.wdk.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabaWdkMarketingItempoolStairQueryitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolStairQueryitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolStairQueryitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolStairQueryitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolStairQueryitemAPIResponseModel is 换购商品查询 成功返回结果
type AlibabaWdkMarketingItempoolStairQueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_stair_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolStairQueryitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolStairQueryitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolStairQueryitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolStairQueryitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolStairQueryitemAPIResponse
func GetAlibabaWdkMarketingItempoolStairQueryitemAPIResponse() *AlibabaWdkMarketingItempoolStairQueryitemAPIResponse {
	return poolAlibabaWdkMarketingItempoolStairQueryitemAPIResponse.Get().(*AlibabaWdkMarketingItempoolStairQueryitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolStairQueryitemAPIResponse 将 AlibabaWdkMarketingItempoolStairQueryitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolStairQueryitemAPIResponse(v *AlibabaWdkMarketingItempoolStairQueryitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolStairQueryitemAPIResponse.Put(v)
}
