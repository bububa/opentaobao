package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票订单查询 APIResponse
taobao.bus.order.get

商家汽车票订单查询
*/
type TaobaoBusOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusOrderGetResponse `json:"bus_order_get_response,omitempty"` 
    TaobaoBusOrderGetResponse
}

/* model for simplify = false
type TaobaoBusOrderGetResponse struct {

    // 订单查询返回对象
    
    Result  *struct {
        B2BOrderQueryRp  *B2BOrderQueryRp `json:"b2b_order_query_rp,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBusOrderGetResponse struct {

    // 订单查询返回对象
    Result   *B2BOrderQueryRp `json:"result,omitempty"`

}
