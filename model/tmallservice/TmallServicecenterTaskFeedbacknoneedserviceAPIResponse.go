package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTaskFeedbacknoneedserviceAPIResponse 服务商反馈无需安装工单接口 API返回值
// tmall.servicecenter.task.feedbacknoneedservice
//
// 服务商反馈无需安装工单接口
type TmallServicecenterTaskFeedbacknoneedserviceAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterTaskFeedbacknoneedserviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel).Reset()
}

// TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel is 服务商反馈无需安装工单接口 成功返回结果
type TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_task_feedbacknoneedservice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterTaskFeedbacknoneedserviceAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterTaskFeedbacknoneedserviceAPIResponse)
	},
}

// GetTmallServicecenterTaskFeedbacknoneedserviceAPIResponse 从 sync.Pool 获取 TmallServicecenterTaskFeedbacknoneedserviceAPIResponse
func GetTmallServicecenterTaskFeedbacknoneedserviceAPIResponse() *TmallServicecenterTaskFeedbacknoneedserviceAPIResponse {
	return poolTmallServicecenterTaskFeedbacknoneedserviceAPIResponse.Get().(*TmallServicecenterTaskFeedbacknoneedserviceAPIResponse)
}

// ReleaseTmallServicecenterTaskFeedbacknoneedserviceAPIResponse 将 TmallServicecenterTaskFeedbacknoneedserviceAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterTaskFeedbacknoneedserviceAPIResponse(v *TmallServicecenterTaskFeedbacknoneedserviceAPIResponse) {
	v.Reset()
	poolTmallServicecenterTaskFeedbacknoneedserviceAPIResponse.Put(v)
}
