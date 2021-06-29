package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全场增加换购品 APIResponse
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品
*/
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingFullrangeAddexchangeitemResponse
}

type AlibabaWdkMarketingFullrangeAddexchangeitemResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_addexchangeitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 添加商品返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
