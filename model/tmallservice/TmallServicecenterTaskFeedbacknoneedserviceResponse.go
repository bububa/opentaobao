package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈无需安装工单接口 APIResponse
tmall.servicecenter.task.feedbacknoneedservice

服务商反馈无需安装工单接口
*/
type TmallServicecenterTaskFeedbacknoneedserviceAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterTaskFeedbacknoneedserviceResponse `json:"tmall_servicecenter_task_feedbacknoneedservice_response,omitempty"`
}

type TmallServicecenterTaskFeedbacknoneedserviceResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}