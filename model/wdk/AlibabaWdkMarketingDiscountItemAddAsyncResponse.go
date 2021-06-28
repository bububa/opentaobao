package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
特价批量新增商品 APIResponse
alibaba.wdk.marketing.discount.item.add.async

新分组模型下新增商品
*/
type AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingDiscountItemAddAsyncResponse
}

type AlibabaWdkMarketingDiscountItemAddAsyncResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_discount_item_add_async_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
