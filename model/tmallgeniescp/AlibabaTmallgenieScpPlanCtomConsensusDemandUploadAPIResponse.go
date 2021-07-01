package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIResponse
22-C2M共识需求回传接口 API返回值
alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload

C2M 共识需求回传接口 */
type AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIResponseModel
}

// AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIResponseModel is 22-C2M共识需求回传接口 成功返回结果
type AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_ctom_consensus_demand_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果msg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 结果code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
