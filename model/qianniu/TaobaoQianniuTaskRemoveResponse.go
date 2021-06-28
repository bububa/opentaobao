package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
轻任务删除接口 APIResponse
taobao.qianniu.task.remove

轻任务删除接口。
*/
type TaobaoQianniuTaskRemoveAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuTaskRemoveResponse `json:"qianniu_task_remove_response,omitempty"` 
    TaobaoQianniuTaskRemoveResponse
}

/* model for simplify = false
type TaobaoQianniuTaskRemoveResponse struct {

    // 是否成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoQianniuTaskRemoveResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
