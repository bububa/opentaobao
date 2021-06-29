package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
15-供应商反馈（原料）同步接口 APIResponse
alibaba.tmallgenie.scp.plan.feedback.raw.upload

供应商反馈（原料）同步接口
*/
type AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanFeedbackRawUploadResponse
}

type AlibabaTmallgenieScpPlanFeedbackRawUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_feedback_raw_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
