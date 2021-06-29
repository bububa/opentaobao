package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过一个码，查询这个码对应的上游企业出库单的单据号 APIResponse
alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode

一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
*/
type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResponse
}

type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_upbillcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
