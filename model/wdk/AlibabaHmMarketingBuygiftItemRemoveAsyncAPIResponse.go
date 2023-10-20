package wdk

import (
	"encoding/xml"

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

// AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel is 批量删除买赠商品 成功返回结果
type AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_buygift_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
