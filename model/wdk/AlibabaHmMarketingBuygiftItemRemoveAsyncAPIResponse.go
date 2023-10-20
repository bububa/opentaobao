package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse 批量删除买赠商品 API返回值
// alibaba.hm.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
type AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel).Reset()
}

// AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel is 批量删除买赠商品 成功返回结果
type AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_buygift_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse)
	},
}

// GetAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse
func GetAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse() *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse {
	return poolAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse.Get().(*AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse)
}

// ReleaseAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse 将 AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse(v *AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse.Put(v)
}
