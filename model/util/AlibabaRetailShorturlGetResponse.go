package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
短链接获取 APIResponse
alibaba.retail.shorturl.get

短链接获取
*/
type AlibabaRetailShorturlGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailShorturlGetResponse `json:"alibaba_retail_shorturl_get_response,omitempty"` 
    AlibabaRetailShorturlGetResponse
}

/* model for simplify = false
type AlibabaRetailShorturlGetResponse struct {

    // result
    
    Result  *struct {
        AlibabaRetailShorturlGetResult  *AlibabaRetailShorturlGetResult `json:"alibaba_retail_shorturl_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailShorturlGetResponse struct {

    // result
    Result   *AlibabaRetailShorturlGetResult `json:"result,omitempty"`

}
