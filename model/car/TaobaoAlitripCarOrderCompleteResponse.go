package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务完成API APIResponse
taobao.alitrip.car.order.complete

用来接收服务商订单流程完成信息
*/
type TaobaoAlitripCarOrderCompleteAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripCarOrderCompleteResponse
}

type TaobaoAlitripCarOrderCompleteResponse struct {
    XMLName xml.Name `xml:"alitrip_car_order_complete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`

    
    // 其它数据
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
