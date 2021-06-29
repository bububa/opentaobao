package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询出入库单处理结果api API返回值 
alibaba.alihealth.tracecodeseller.bill.result.search

查询出入库单处理结果api
*/
type AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerBillResultSearchResponse
}

// 查询出入库单处理结果api 成功返回结果
type AlibabaAlihealthTracecodesellerBillResultSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_bill_result_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
