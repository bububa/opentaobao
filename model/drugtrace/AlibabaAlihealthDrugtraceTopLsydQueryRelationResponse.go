package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询 APIResponse
alibaba.alihealth.drugtrace.top.lsyd.query.relation

单码关联关系查询
*/
type AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopLsydQueryRelationResponse
}

type AlibabaAlihealthDrugtraceTopLsydQueryRelationResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_relation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugtraceTopLsydQueryRelationResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
