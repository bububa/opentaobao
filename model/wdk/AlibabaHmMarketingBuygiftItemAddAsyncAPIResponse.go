package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse 批量发布买赠商品 API返回值
// alibaba.hm.marketing.buygift.item.add.async
//
// 批量发布买赠商品
type AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingBuygiftItemAddAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingBuygiftItemAddAsyncAPIResponseModel).Reset()
}

// AlibabaHmMarketingBuygiftItemAddAsyncAPIResponseModel is 批量发布买赠商品 成功返回结果
type AlibabaHmMarketingBuygiftItemAddAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_buygift_item_add_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingBuygiftItemAddAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse)
	},
}

// GetAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse
func GetAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse() *AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse {
	return poolAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse.Get().(*AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse)
}

// ReleaseAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse 将 AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse(v *AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingBuygiftItemAddAsyncAPIResponse.Put(v)
}
