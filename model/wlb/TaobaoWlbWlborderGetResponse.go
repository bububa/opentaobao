package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝订单编号查询物流宝订单概要信息 APIResponse
taobao.wlb.wlborder.get

根据物流宝订单编号查询物流宝订单概要信息
*/
type TaobaoWlbWlborderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWlborderGetResponse `json:"wlb_wlborder_get_response,omitempty"` 
    TaobaoWlbWlborderGetResponse
}

/* model for simplify = false
type TaobaoWlbWlborderGetResponse struct {

    // 物流宝订单对象
    
    WlbOrder  *struct {
        WlbOrder  *WlbOrder `json:"wlb_order,omitempty"`
    } `json:"wlb_order,omitempty"`
    

}
*/

type TaobaoWlbWlborderGetResponse struct {

    // 物流宝订单对象
    WlbOrder   *WlbOrder `json:"wlb_order,omitempty"`

}
