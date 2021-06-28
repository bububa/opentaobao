package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈无需安装工单接口 APIResponse
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口
*/
type TmallServicecenterTaskFeedbacknoneedserviceAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_task_feedbacknoneedservice_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"