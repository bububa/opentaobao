package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询物流宝订单 APIResponse
taobao.wlb.order.page.get

分页查询物流宝订单
*/
type TaobaoWlbOrderPageGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbOrderPageGetResponse `json:"wlb_order_page_get_response,omitempty"` 
    TaobaoWlbOrderPageGetResponse
}

/* model for simplify = false
type TaobaoWlbOrderPageGetResponse struct {

    // 物流宝订单对象
    
    OrderList  struct {
        WlbOrder  []WlbOrder `json:"wlb_order,omitempty"`
    } `json:"order_list,omitempty"`
    

    // 总条数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

type TaobaoWlbOrderPageGetResponse struct {

    // 物流宝订单对象
    OrderList   []WlbOrder `json:"order_list,omitempty"`

    // 总条数
    TotalCount   int64 `json:"total_count,omitempty"`

}
