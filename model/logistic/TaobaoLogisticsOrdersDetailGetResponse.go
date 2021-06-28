package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单,返回详细信息 APIResponse
taobao.logistics.orders.detail.get

查询物流订单的详细信息，涉及用户隐私字段。
*/
type TaobaoLogisticsOrdersDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsOrdersDetailGetResponse `json:"logistics_orders_detail_get_response,omitempty"` 
    TaobaoLogisticsOrdersDetailGetResponse
}

/* model for simplify = false
type TaobaoLogisticsOrdersDetailGetResponse struct {

    // 获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息.
    
    Shippings  struct {
        Shipping  []Shipping `json:"shipping,omitempty"`
    } `json:"shippings,omitempty"`
    

    // 搜索到的物流订单列表总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoLogisticsOrdersDetailGetResponse struct {

    // 获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息.
    Shippings   []Shipping `json:"shippings,omitempty"`

    // 搜索到的物流订单列表总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
