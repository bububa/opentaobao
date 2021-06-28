package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
任务查询条数接口 APIResponse
taobao.qianniu.tasks.count

任务查询条数接口
*/
type TaobaoQianniuTasksCountAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuTasksCountResponse `json:"qianniu_tasks_count_response,omitempty"` 
    TaobaoQianniuTasksCountResponse
}

/* model for simplify = false
type TaobaoQianniuTasksCountResponse struct {

    // 符合查询条件的总条数
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type TaobaoQianniuTasksCountResponse struct {

    // 符合查询条件的总条数
    Result   int64 `json:"result,omitempty"`

}
