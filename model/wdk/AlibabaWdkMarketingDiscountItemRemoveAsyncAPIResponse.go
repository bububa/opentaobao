package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse
特价批量移除商品 API返回值
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品 */
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel
}

// AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel is 特价批量移除商品 成功返回结果
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_discount_item_remove_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
