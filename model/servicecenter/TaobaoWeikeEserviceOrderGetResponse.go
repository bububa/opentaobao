package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单查询 APIResponse
taobao.weike.eservice.order.get

用于客服外包中服务商查询订单列表
*/
type TaobaoWeikeEserviceOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWeikeEserviceOrderGetResponse `json:"weike_eservice_order_get_response,omitempty"` 
    TaobaoWeikeEserviceOrderGetResponse
}

/* model for simplify = false
type TaobaoWeikeEserviceOrderGetResponse struct {

    // 记录总记录数
    
    TotalSize   int64 `json:"total_size,omitempty"`
    

    // 订单列表
    
    OrderList  struct {
        Order  []Order `json:"order,omitempty"`
    } `json:"order_list,omitempty"`
    

}
*/

type TaobaoWeikeEserviceOrderGetResponse struct {

    // 记录总记录数
    TotalSize   int64 `json:"total_size,omitempty"`

    // 订单列表
    OrderList   []Order `json:"order_list,omitempty"`

}
