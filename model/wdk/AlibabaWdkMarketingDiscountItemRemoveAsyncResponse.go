package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
特价批量移除商品 APIResponse
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品
*/
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_discount_item_remove_async_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *MarketResult `json:"result,omitempty" xml:"