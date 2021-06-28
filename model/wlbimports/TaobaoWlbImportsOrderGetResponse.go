package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流订单获取 APIResponse
taobao.wlb.imports.order.get

一般进口物流订单获取
*/
type TaobaoWlbImportsOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbImportsOrderGetResponse `json:"wlb_imports_order_get_response,omitempty"` 
    TaobaoWlbImportsOrderGetResponse
}

/* model for simplify = false
type TaobaoWlbImportsOrderGetResponse struct {

    // 物流订单信息
    
    Orders  struct {
        LocOrder  []LocOrder `json:"loc_order,omitempty"`
    } `json:"orders,omitempty"`
    

    // 搜索到的总数量
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoWlbImportsOrderGetResponse struct {

    // 物流订单信息
    Orders   []LocOrder `json:"orders,omitempty"`

    // 搜索到的总数量
    TotalResults   int64 `json:"total_results,omitempty"`

}
