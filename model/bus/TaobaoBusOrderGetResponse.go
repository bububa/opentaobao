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
    Response *TaobaoBusOrderGetResponse `json:"taobao_bus_order_get_response,omitempty"`
}

type TaobaoBusOrderGetResponse struct {

    // 订单查询返回对象
    Result   *B2BOrderQueryRp `json:"result,omitempty"`

}
