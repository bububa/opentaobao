package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolCreateactivityAPIResponse 添加商品池活动 API返回值
// alibaba.wdk.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabaWdkMarketingItempoolCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolCreateactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolCreateactivityAPIResponseModel is 添加商品池活动 成功返回结果
type AlibabaWdkMarketingItempoolCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolCreateactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolCreateactivityAPIResponse
func GetAlibabaWdkMarketingItempoolCreateactivityAPIResponse() *AlibabaWdkMarketingItempoolCreateactivityAPIResponse {
	return poolAlibabaWdkMarketingItempoolCreateactivityAPIResponse.Get().(*AlibabaWdkMarketingItempoolCreateactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolCreateactivityAPIResponse 将 AlibabaWdkMarketingItempoolCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolCreateactivityAPIResponse(v *AlibabaWdkMarketingItempoolCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolCreateactivityAPIResponse.Put(v)
}
