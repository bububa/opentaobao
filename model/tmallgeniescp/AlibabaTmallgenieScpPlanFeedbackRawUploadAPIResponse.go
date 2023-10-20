package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse 15-供应商反馈（原料）同步接口 API返回值
// alibaba.tmallgenie.scp.plan.feedback.raw.upload
//
// 供应商反馈（原料）同步接口
type AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponseModel is 15-供应商反馈（原料）同步接口 成功返回结果
type AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_feedback_raw_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse
func GetAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse() *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse 将 AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse(v *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse.Put(v)
}
