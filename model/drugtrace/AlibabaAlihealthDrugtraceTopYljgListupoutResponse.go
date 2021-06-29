package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
医疗机构查询本企业上游企业出库单据信息 APIResponse
alibaba.alihealth.drugtrace.top.yljg.listupout

查询货主/本企业上游企业出库单据信息
*/
type AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopYljgListupoutResponse
}

type AlibabaAlihealthDrugtraceTopYljgListupoutResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_listupout_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugtraceTopYljgListupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
