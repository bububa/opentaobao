package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
生成发布使用的版本号 APIResponse
alibaba.wdk.marketing.version.generate

生成发布使用的版本号
*/
type AlibabaWdkMarketingVersionGenerateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingVersionGenerateResponse `json:"alibaba_wdk_marketing_version_generate_response,omitempty"` 
    AlibabaWdkMarketingVersionGenerateResponse
}

/* model for simplify = false
type AlibabaWdkMarketingVersionGenerateResponse struct {

    // result
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingVersionGenerateResponse struct {

    // result
    Result   *MarketResult `json:"result,omitempty"`

}
