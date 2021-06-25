package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
任务列表获取接口 APIResponse
taobao.vmarket.eticket.tasks.get

外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号
*/
type TaobaoVmarketEticketTasksGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketTasksGetResponse `json:"taobao_vmarket_eticket_tasks_get_response,omitempty"`
}

type TaobaoVmarketEticketTasksGetResponse struct {

    // 任务列表查询结果的总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 任务列表查询结果信息
    EticketTasks   []EticketTask `json:"eticket_tasks,omitempty"`

}
