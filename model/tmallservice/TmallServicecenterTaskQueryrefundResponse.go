package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单是否退款 APIResponse
tmall.servicecenter.task.queryrefund

查询任务类工单是否退款
*/
type TmallServicecenterTaskQueryrefundAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterTaskQueryrefundResponse `json:"tmall_servicecenter_task_queryrefund_response,omitempty"`
}

type TmallServicecenterTaskQueryrefundResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
