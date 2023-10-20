package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse 特价批量移除商品 API返回值
// alibaba.wdk.marketing.discount.item.remove.async
//
// 特价批量移除商品
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel is 特价批量移除商品 成功返回结果
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_discount_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse
func GetAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse() *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse {
	return poolAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse.Get().(*AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse 将 AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse(v *AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse.Put(v)
}
