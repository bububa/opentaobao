package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingPriceAPIResponse 促销价签服务 API返回值
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
type AlibabaWdkMarketingPriceAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingPriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingPriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingPriceAPIResponseModel).Reset()
}

// AlibabaWdkMarketingPriceAPIResponseModel is 促销价签服务 成功返回结果
type AlibabaWdkMarketingPriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_price_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PromotionPriceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingPriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingPriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingPriceAPIResponse)
	},
}

// GetAlibabaWdkMarketingPriceAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingPriceAPIResponse
func GetAlibabaWdkMarketingPriceAPIResponse() *AlibabaWdkMarketingPriceAPIResponse {
	return poolAlibabaWdkMarketingPriceAPIResponse.Get().(*AlibabaWdkMarketingPriceAPIResponse)
}

// ReleaseAlibabaWdkMarketingPriceAPIResponse 将 AlibabaWdkMarketingPriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingPriceAPIResponse(v *AlibabaWdkMarketingPriceAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingPriceAPIResponse.Put(v)
}
