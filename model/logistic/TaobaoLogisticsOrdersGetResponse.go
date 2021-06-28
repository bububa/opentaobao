package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单 APIResponse
taobao.logistics.orders.get

批量查询物流订单。
*/
type TaobaoLogisticsOrdersGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsOrdersGetResponse `json:"logistics_orders_get_response,omitempty"` 
    TaobaoLogisticsOrdersGetResponse
}

/* model for simplify = false
type TaobaoLogisticsOrdersGetResponse struct {

    // 获取的物流订单详情列表 <br/>返回的Shipping包含的具体信息为入参fields请求的字段信息
    
    Shippings  struct {
        Shipping  []Shipping `json:"shipping,omitempty"`
    } `json:"shippings,omitempty"`
    

    // 搜索到的物流订单列表总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoLogisticsOrdersGetResponse struct {

    // 获取的物流订单详情列表 <br/>返回的Shipping包含的具体信息为入参fields请求的字段信息
    Shippings   []Shipping `json:"shippings,omitempty"`

    // 搜索到的物流订单列表总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
