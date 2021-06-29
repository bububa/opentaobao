package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生成发布使用的版本号 APIResponse
alibaba.wdk.marketing.version.generate

生成发布使用的版本号
*/
type AlibabaWdkMarketingVersionGenerateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingVersionGenerateResponse
}

type AlibabaWdkMarketingVersionGenerateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_version_generate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
