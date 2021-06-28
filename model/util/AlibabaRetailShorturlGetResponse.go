package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短链接获取 APIResponse
alibaba.retail.shorturl.get

短链接获取
*/
type AlibabaRetailShorturlGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_retail_shorturl_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaRetailShorturlGetResult `json:"result,omitempty" xml:"