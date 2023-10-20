package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolAdditemAPIResponse 增加商品池里面的商品 API返回值
// alibaba.wdk.marketing.itempool.additem
//
// 增加商品池里面的商品
type AlibabaWdkMarketingItempoolAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolAdditemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolAdditemAPIResponseModel is 增加商品池里面的商品 成功返回结果
type AlibabaWdkMarketingItempoolAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolAdditemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolAdditemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolAdditemAPIResponse
func GetAlibabaWdkMarketingItempoolAdditemAPIResponse() *AlibabaWdkMarketingItempoolAdditemAPIResponse {
	return poolAlibabaWdkMarketingItempoolAdditemAPIResponse.Get().(*AlibabaWdkMarketingItempoolAdditemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolAdditemAPIResponse 将 AlibabaWdkMarketingItempoolAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolAdditemAPIResponse(v *AlibabaWdkMarketingItempoolAdditemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolAdditemAPIResponse.Put(v)
}
