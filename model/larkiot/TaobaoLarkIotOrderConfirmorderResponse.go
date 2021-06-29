package larkiot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道卖品落单 APIResponse
taobao.lark.iot.order.confirmorder

云智对接无人超市，接收无人超市订单信息
*/
type TaobaoLarkIotOrderConfirmorderAPIResponse struct {
    model.CommonResponse
    TaobaoLarkIotOrderConfirmorderResponse
}

type TaobaoLarkIotOrderConfirmorderResponse struct {
    XMLName xml.Name `xml:"lark_iot_order_confirmorder_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回下单结果
    
    Data   *BizSingleResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
