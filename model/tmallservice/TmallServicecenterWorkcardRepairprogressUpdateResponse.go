package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新维修进度 APIResponse
tmall.servicecenter.workcard.repairprogress.update

提供给外部合作服务商的维修进度更改接口
*/
type TmallServicecenterWorkcardRepairprogressUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardRepairprogressUpdateResponse `json:"tmall_servicecenter_workcard_repairprogress_update_response,omitempty"`
}

type TmallServicecenterWorkcardRepairprogressUpdateResponse struct {

    // 返回结果
    Result   *TmallServicecenterWorkcardRepairprogressUpdateResult `json:"result,omitempty"`

}
