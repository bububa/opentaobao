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
    Response *AlibabaRetailShorturlGetResponse `json:"alibaba_retail_shorturl_get_response,omitempty"`
}

type AlibabaRetailShorturlGetResponse struct {

    // result
    Result   *AlibabaRetailShorturlGetResult `json:"result,omitempty"`

}
