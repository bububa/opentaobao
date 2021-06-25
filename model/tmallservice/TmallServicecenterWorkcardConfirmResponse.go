package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商确认服务完成 APIResponse
tmall.servicecenter.workcard.confirm

提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
*/
type TmallServicecenterWorkcardConfirmAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardConfirmResponse `json:"tmall_servicecenter_workcard_confirm_response,omitempty"`
}

type TmallServicecenterWorkcardConfirmResponse struct {

    // 请求结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
