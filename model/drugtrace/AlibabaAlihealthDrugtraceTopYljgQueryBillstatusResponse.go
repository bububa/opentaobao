package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单据后处理状态查询 API返回值 
alibaba.alihealth.drugtrace.top.yljg.query.billstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResponse
}

// 上传单据后处理状态查询 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_billstatus_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
