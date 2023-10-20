package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse 特价批量新增商品 API返回值
// alibaba.wdk.marketing.discount.item.add.async
//
// 新分组模型下新增商品
type AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingDiscountItemAddAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingDiscountItemAddAsyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingDiscountItemAddAsyncAPIResponseModel is 特价批量新增商品 成功返回结果
type AlibabaWdkMarketingDiscountItemAddAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_discount_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingDiscountItemAddAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse
func GetAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse() *AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse {
	return poolAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse.Get().(*AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse 将 AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse(v *AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingDiscountItemAddAsyncAPIResponse.Put(v)
}
