package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanNetdemandUploadAPIResponse 23-Net Demand净需求回传 API返回值
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
type AlibabaTmallgenieScpPlanNetdemandUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanNetdemandUploadAPIResponseModel
}

// AlibabaTmallgenieScpPlanNetdemandUploadAPIResponseModel is 23-Net Demand净需求回传 成功返回结果
type AlibabaTmallgenieScpPlanNetdemandUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_netdemand_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
