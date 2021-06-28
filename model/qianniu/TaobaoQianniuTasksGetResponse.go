package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定的任务 APIResponse
taobao.qianniu.tasks.get

获取指定的任务，可用的参数组合：<br/>task_ids + need_meta + fields：精确查找<br/>biz_type + sub_biz_type + biz_ids + need_meta + fields：按照业务ID查找<br/>biz_type + sub_biz_type + sender_uid + need_meta + fields：按照发起者查找<br/>biz_type + sub_biz_type + receiver_uid + need_meta + fields：按照执行者查找<br/>biz_type+modify_start_time+modify_end_time+fields:能支持指定修改时间的查询，用于增量查询等
*/
type TaobaoQianniuTasksGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuTasksGetResponse `json:"qianniu_tasks_get_response,omitempty"` 
    TaobaoQianniuTasksGetResponse
}

/* model for simplify = false
type TaobaoQianniuTasksGetResponse struct {

    // 返回的任务列表
    
    Tasks  struct {
        QTask  []QTask `json:"q_task,omitempty"`
    } `json:"tasks,omitempty"`
    

}
*/

type TaobaoQianniuTasksGetResponse struct {

    // 返回的任务列表
    Tasks   []QTask `json:"tasks,omitempty"`

}
