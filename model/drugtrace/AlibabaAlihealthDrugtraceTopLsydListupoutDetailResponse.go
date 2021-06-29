package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上游出库单单据明细查询 APIResponse
alibaba.alihealth.drugtrace.top.lsyd.listupout.detail

查询上游出库单明细(带追溯码信息)
*/
type AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugtraceTopLsydListupoutDetailResponse
}

type AlibabaAlihealthDrugtraceTopLsydListupoutDetailResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_listupout_detail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugtraceTopLsydListupoutDetailResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
