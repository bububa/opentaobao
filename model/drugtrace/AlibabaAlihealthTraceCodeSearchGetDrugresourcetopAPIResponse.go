package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取码信息 API返回值 
alibaba.alihealth.trace.code.search.get.drugresourcetop

根据码获取码信息
*/
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponseModel
}

// 根据码获取码信息 成功返回结果
type AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_trace_code_search_get_drugresourcetop_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
