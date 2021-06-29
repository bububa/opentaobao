package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单创建 APIResponse
taobao.opentrade.create.order

交易开放创建订单
*/
type TaobaoOpentradeCreateOrderAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeCreateOrderResponse
}

type TaobaoOpentradeCreateOrderResponse struct {
    XMLName xml.Name `xml:"opentrade_create_order_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单ID
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`

    
}
