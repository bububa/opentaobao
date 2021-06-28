package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流宝订单流转状态查询 APIResponse
taobao.wlb.orderstatus.get

根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
*/
type TaobaoWlbOrderstatusGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbOrderstatusGetResponse `json:"wlb_orderstatus_get_response,omitempty"` 
    TaobaoWlbOrderstatusGetResponse
}

/* model for simplify = false
type TaobaoWlbOrderstatusGetResponse struct {

    // 订单流转信息状态列表
    
    WlbOrderStatus  struct {
        WlbProcessStatus  []WlbProcessStatus `json:"wlb_process_status,omitempty"`
    } `json:"wlb_order_status,omitempty"`
    

}
*/

type TaobaoWlbOrderstatusGetResponse struct {

    // 订单流转信息状态列表
    WlbOrderStatus   []WlbProcessStatus `json:"wlb_order_status,omitempty"`

}
