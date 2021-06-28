package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务列表获取接口 APIResponse
taobao.vmarket.eticket.tasks.get

外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号
*/
type TaobaoVmarketEticketTasksGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_tasks_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 任务列表查询结果的总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"