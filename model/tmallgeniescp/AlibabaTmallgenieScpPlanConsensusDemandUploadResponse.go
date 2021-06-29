package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
20-IBP共识需求回传接口 APIResponse
alibaba.tmallgenie.scp.plan.consensus.demand.upload

IBP共识需求回传接口
*/
type AlibabaTmallgenieScpPlanConsensusDemandUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanConsensusDemandUploadResponse
}

type AlibabaTmallgenieScpPlanConsensusDemandUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_consensus_demand_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaTmallgenieScpPlanConsensusDemandUploadResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
