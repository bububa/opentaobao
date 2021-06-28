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
    // Response *TaobaoQianniuTaskFinishResponse `json:"qianniu_task_finish_response,omitempty"` 
    TaobaoQianniuTaskFinishResponse
}

/* model for simplify = false
type TaobaoQianniuTaskFinishResponse struct {

    // 是否成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoQianniuTaskFinishResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
