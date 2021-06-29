package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单据后处理状态查询 APIResponse
alibaba.alihealth.drugtrace.top.lsyd.query.billstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResponse
}

type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_billstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
