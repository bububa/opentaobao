package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票下单接口 APIResponse
taobao.bus.order.set

提供给汽车票商家进行下单
*/
type TaobaoBusOrderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_order_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 支付宝交易流水号
    
    AliPayTradeId   string `json:"ali_pay_trade_id,omitempty" xml:"