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
    AlibabaRetailShorturlGetResponse
}

type AlibabaRetailShorturlGetResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_shorturl_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaRetailShorturlGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
