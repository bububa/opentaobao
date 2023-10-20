package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse 批量发布买赠商品 API返回值
// alibaba.wdk.marketing.buygift.item.add.async
//
// 批量发布买赠商品
type AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponseModel is 批量发布买赠商品 成功返回结果
type AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_buygift_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse
func GetAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse() *AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse {
	return poolAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse.Get().(*AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse 将 AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse(v *AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse.Put(v)
}
