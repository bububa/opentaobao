package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商api APIResponse
alibaba.alihealth.tracecodeseller.channel.search

查询渠道商api
*/
type AlibabaAlihealthTracecodesellerChannelSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerChannelSearchResponse
}

type AlibabaAlihealthTracecodesellerChannelSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_channel_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
