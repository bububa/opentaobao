package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
促销价签服务 APIResponse
alibaba.wdk.marketing.price

获取营销-促销商品中的实时价格
*/
type AlibabaWdkMarketingPriceAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingPriceResponse
}

type AlibabaWdkMarketingPriceResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_price_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *PromotionPriceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
