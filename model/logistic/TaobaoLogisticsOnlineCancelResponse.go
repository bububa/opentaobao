package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消物流订单接口 APIResponse
taobao.logistics.online.cancel

调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。
*/
type TaobaoLogisticsOnlineCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsOnlineCancelResponse `json:"logistics_online_cancel_response,omitempty"` 
    TaobaoLogisticsOnlineCancelResponse
}

/* model for simplify = false
type TaobaoLogisticsOnlineCancelResponse struct {

    // 成功与失败
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 修改时间
    
    ModifyTime   string `json:"modify_time,omitempty"`
    

    // 重新创建物流订单id
    
    RecreatedOrderId   string `json:"recreated_order_id,omitempty"`
    

}
*/

type TaobaoLogisticsOnlineCancelResponse struct {

    // 成功与失败
    IsSuccess   bool `json:"is_success,omitempty"`

    // 修改时间
    ModifyTime   string `json:"modify_time,omitempty"`

    // 重新创建物流订单id
    RecreatedOrderId   string `json:"recreated_order_id,omitempty"`

}
