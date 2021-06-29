package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
配送拦截接口 APIResponse
taobao.qimen.order.callback

配送拦截
*/
type TaobaoQimenOrderCallbackAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderCallbackResponse
}

type TaobaoQimenOrderCallbackResponse struct {
    XMLName xml.Name `xml:"qimen_order_callback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *OrderCallbackResponseDO `json:"response,omitempty" xml:"response,omitempty"`

    
}
