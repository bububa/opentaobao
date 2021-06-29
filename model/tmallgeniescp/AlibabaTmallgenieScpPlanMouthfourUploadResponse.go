package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
21-M+4PR 回传接口接口 APIResponse
alibaba.tmallgenie.scp.plan.mouthfour.upload

M+4 PR 回传接口
*/
type AlibabaTmallgenieScpPlanMouthfourUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanMouthfourUploadResponse
}

type AlibabaTmallgenieScpPlanMouthfourUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_mouthfour_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果msg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 请求唯一ID
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 结果code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
