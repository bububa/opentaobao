package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料-LT内的POGAP数据回传 APIResponse
alibaba.tmallgenie.scp.plan.rawpo.gap.return

二级物料-LT内的POGAP数据回传
*/
type AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanRawpoGapReturnResponse
}

type AlibabaTmallgenieScpPlanRawpoGapReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_rawpo_gap_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 请求唯一ID
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 返回码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
