package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家信息 APIResponse
alibaba.alihealth.tracecodeseller.ent.search

查询商家信息
*/
type AlibabaAlihealthTracecodesellerEntSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerEntSearchResponse
}

type AlibabaAlihealthTracecodesellerEntSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_ent_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
