package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingDiscountItemAddAsyncAPIResponse 特价批量新增商品 API返回值
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
type AlibabaHmMarketingDiscountItemAddAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingDiscountItemAddAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingDiscountItemAddAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingDiscountItemAddAsyncAPIResponseModel).Reset()
}

// AlibabaHmMarketingDiscountItemAddAsyncAPIResponseModel is 特价批量新增商品 成功返回结果
type AlibabaHmMarketingDiscountItemAddAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_discount_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingDiscountItemAddAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingDiscountItemAddAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingDiscountItemAddAsyncAPIResponse)
	},
}

// GetAlibabaHmMarketingDiscountItemAddAsyncAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingDiscountItemAddAsyncAPIResponse
func GetAlibabaHmMarketingDiscountItemAddAsyncAPIResponse() *AlibabaHmMarketingDiscountItemAddAsyncAPIResponse {
	return poolAlibabaHmMarketingDiscountItemAddAsyncAPIResponse.Get().(*AlibabaHmMarketingDiscountItemAddAsyncAPIResponse)
}

// ReleaseAlibabaHmMarketingDiscountItemAddAsyncAPIResponse 将 AlibabaHmMarketingDiscountItemAddAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingDiscountItemAddAsyncAPIResponse(v *AlibabaHmMarketingDiscountItemAddAsyncAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingDiscountItemAddAsyncAPIResponse.Put(v)
}
