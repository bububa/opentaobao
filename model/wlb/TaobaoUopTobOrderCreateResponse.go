package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ToB仓储发货 APIResponse
taobao.uop.tob.order.create

ToB仓储发货
*/
type TaobaoUopTobOrderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoUopTobOrderCreateResponse
}

type TaobaoUopTobOrderCreateResponse struct {
    XMLName xml.Name `xml:"uop_tob_order_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // flag
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 订单
    
    DeliveryOrders   []Deliveryorder `json:"delivery_orders,omitempty" xml:"delivery_orders>deliveryorder,omitempty"`
    
    
}
