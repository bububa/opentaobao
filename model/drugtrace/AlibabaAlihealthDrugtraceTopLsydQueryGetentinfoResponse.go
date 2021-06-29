package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过企业名得到唯一标识ref_ent_id及企业ent_id API返回值 
alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo

根据企业名称查询ID
*/
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResponse
}

// 通过企业名得到唯一标识ref_ent_id及企业ent_id 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_getentinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 监控宝推送网站监控信息，返回结果
    Result   *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
