package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取码信息 APIResponse
alibaba.alihealth.trace.code.search.get.drugresourcetop

根据码获取码信息
*/
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTraceCodeSearchGetDrugresourcetopResponse
}

type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_trace_code_search_get_drugresourcetop_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
