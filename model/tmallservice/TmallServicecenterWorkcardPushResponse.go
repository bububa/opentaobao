package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推送服务工单信息 APIResponse
tmall.servicecenter.workcard.push

服务商家推送工单信息到天猫。
*/
type TmallServicecenterWorkcardPushAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardPushResponse `json:"tmall_servicecenter_workcard_push_response,omitempty"`
}

type TmallServicecenterWorkcardPushResponse struct {

    // 返回结果
    WorkcardResult   *ResultBase `json:"workcard_result,omitempty"`

}
