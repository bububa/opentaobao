package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse 14-供应商反馈（OEM）同步接口 API返回值
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
type AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponseModel is 14-供应商反馈（OEM）同步接口 成功返回结果
type AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_feedback_oem_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse
func GetAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse() *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse 将 AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse(v *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse.Put(v)
}
