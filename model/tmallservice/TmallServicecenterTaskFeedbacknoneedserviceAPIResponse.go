package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTaskFeedbacknoneedserviceAPIResponse
服务商反馈无需安装工单接口 API返回值
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口 */
type TmallServicecenterTaskFeedbacknoneedserviceAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel
}

// TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel is 服务商反馈无需安装工单接口 成功返回结果
type TmallServicecenterTaskFeedbacknoneedserviceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_task_feedbacknoneedservice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
