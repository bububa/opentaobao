package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolQueryactivityAPIResponse 查找商品池活动 API返回值
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
type AlibabaWdkMarketingItempoolQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolQueryactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolQueryactivityAPIResponseModel is 查找商品池活动 成功返回结果
type AlibabaWdkMarketingItempoolQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolQueryactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolQueryactivityAPIResponse
func GetAlibabaWdkMarketingItempoolQueryactivityAPIResponse() *AlibabaWdkMarketingItempoolQueryactivityAPIResponse {
	return poolAlibabaWdkMarketingItempoolQueryactivityAPIResponse.Get().(*AlibabaWdkMarketingItempoolQueryactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolQueryactivityAPIResponse 将 AlibabaWdkMarketingItempoolQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolQueryactivityAPIResponse(v *AlibabaWdkMarketingItempoolQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolQueryactivityAPIResponse.Put(v)
}
