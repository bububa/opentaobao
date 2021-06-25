package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
完成轻任务 APIResponse
taobao.qianniu.task.finish

由任务执行者调用
*/
type TaobaoQianniuTaskFinishAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQianniuTaskFinishResponse `json:"taobao_qianniu_task_finish_response,omitempty"`
}

type TaobaoQianniuTaskFinishResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
