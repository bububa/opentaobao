package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消轻任务 APIResponse
taobao.qianniu.task.cancel

由任务发起者调用
*/
type TaobaoQianniuTaskCancelAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQianniuTaskCancelResponse `json:"taobao_qianniu_task_cancel_response,omitempty"`
}

type TaobaoQianniuTaskCancelResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
