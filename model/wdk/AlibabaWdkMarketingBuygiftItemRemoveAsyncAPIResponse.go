package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse 批量删除买赠商品 API返回值
// alibaba.wdk.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
type AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponseModel is 批量删除买赠商品 成功返回结果
type AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_buygift_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse
func GetAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse() *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse {
	return poolAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse.Get().(*AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse 将 AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse(v *AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse.Put(v)
}
