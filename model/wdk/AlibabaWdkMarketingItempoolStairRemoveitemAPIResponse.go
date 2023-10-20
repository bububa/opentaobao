package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse 删除换购活动商品 API返回值
// alibaba.wdk.marketing.itempool.stair.removeitem
//
// 删除换购商品
type AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolStairRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolStairRemoveitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolStairRemoveitemAPIResponseModel is 删除换购活动商品 成功返回结果
type AlibabaWdkMarketingItempoolStairRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_stair_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolStairRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse
func GetAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse() *AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse {
	return poolAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse.Get().(*AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse 将 AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse(v *AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolStairRemoveitemAPIResponse.Put(v)
}
