package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ToB仓储发货 API返回值 
taobao.uop.tob.order.create

ToB仓储发货
*/
type TaobaoUopTobOrderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoUopTobOrderCreateResponse
}

// ToB仓储发货 成功返回结果
type TaobaoUopTobOrderCreateResponse struct {
    XMLName xml.Name `xml:"uop_tob_order_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // flag
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 订单
    DeliveryOrders   []DeliveryOrder `json:"delivery_orders,omitempty" xml:"delivery_orders>delivery_order,omitempty"`
}
