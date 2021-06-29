package alitripcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
司机或客服取消订单 APIResponse
taobao.alitrip.car.order.agent.cancel

司机或客服取消订单后通知飞猪订单取消
*/
type TaobaoAlitripCarOrderAgentCancelAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripCarOrderAgentCancelResponse
}

type TaobaoAlitripCarOrderAgentCancelResponse struct {
    XMLName xml.Name `xml:"alitrip_car_order_agent_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`

    
}
