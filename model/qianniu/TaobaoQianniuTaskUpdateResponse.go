package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新轻任务 APIResponse
taobao.qianniu.task.update

由任务执行者调用，sub_status，tag和memo至少提供一个
*/
type TaobaoQianniuTaskUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQianniuTaskUpdateResponse `json:"taobao_qianniu_task_update_response,omitempty"`
}

type TaobaoQianniuTaskUpdateResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
