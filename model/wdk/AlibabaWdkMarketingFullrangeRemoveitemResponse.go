package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全场活动删除购品 APIResponse
alibaba.wdk.marketing.fullrange.removeitem

删除换购商品
*/
type AlibabaWdkMarketingFullrangeRemoveitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingFullrangeRemoveitemResponse
}

type AlibabaWdkMarketingFullrangeRemoveitemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_removeitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
