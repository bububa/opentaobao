package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
配送拦截接口 API返回值 
taobao.qimen.order.callback

配送拦截
*/
type TaobaoQimenOrderCallbackAPIResponse struct {
    model.CommonResponse
    TaobaoQimenOrderCallbackResponse
}

// 配送拦截接口 成功返回结果
type TaobaoQimenOrderCallbackResponse struct {
    XMLName xml.Name `xml:"qimen_order_callback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *OrderCallbackResponseDO `json:"response,omitempty" xml:"response,omitempty"`
}
