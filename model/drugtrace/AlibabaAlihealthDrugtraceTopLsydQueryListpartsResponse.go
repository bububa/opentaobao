package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
往来单位查询 APIResponse
alibaba.alihealth.drugtrace.top.lsyd.query.listparts

查询往来单位列表
*/
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopLsydQueryListpartsResponse
}

type AlibabaAlihealthDrugtraceTopLsydQueryListpartsResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_listparts_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugtraceTopLsydQueryListpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
