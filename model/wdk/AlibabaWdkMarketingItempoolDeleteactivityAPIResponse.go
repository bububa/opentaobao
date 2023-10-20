package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolDeleteactivityAPIResponse 删除商品池活动 API返回值
// alibaba.wdk.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabaWdkMarketingItempoolDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolDeleteactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolDeleteactivityAPIResponseModel is 删除商品池活动 成功返回结果
type AlibabaWdkMarketingItempoolDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolDeleteactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolDeleteactivityAPIResponse
func GetAlibabaWdkMarketingItempoolDeleteactivityAPIResponse() *AlibabaWdkMarketingItempoolDeleteactivityAPIResponse {
	return poolAlibabaWdkMarketingItempoolDeleteactivityAPIResponse.Get().(*AlibabaWdkMarketingItempoolDeleteactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolDeleteactivityAPIResponse 将 AlibabaWdkMarketingItempoolDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolDeleteactivityAPIResponse(v *AlibabaWdkMarketingItempoolDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolDeleteactivityAPIResponse.Put(v)
}
