package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发布买赠商品 APIResponse
alibaba.wdk.marketing.buygift.item.add.async

批量发布买赠商品
*/
type AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingBuygiftItemAddAsyncResponse
}

type AlibabaWdkMarketingBuygiftItemAddAsyncResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_buygift_item_add_async_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
