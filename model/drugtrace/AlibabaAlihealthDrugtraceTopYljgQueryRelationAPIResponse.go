package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询 API返回值 
alibaba.alihealth.drugtrace.top.yljg.query.relation

单码关联关系查询
*/
type AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel
}

// 单码关联关系查询 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgQueryRelationAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_relation_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihealthDrugtraceTopYljgQueryRelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
