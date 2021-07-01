package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商api API返回值 
alibaba.alihealth.tracecodeseller.channel.search

查询渠道商api
*/
type AlibabaAlihealthTracecodesellerChannelSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel
}

// 查询渠道商api 成功返回结果
type AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_channel_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
