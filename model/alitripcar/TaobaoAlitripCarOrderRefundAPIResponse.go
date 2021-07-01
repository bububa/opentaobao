package alitripcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户投诉达成一致后给用户退款 API返回值 
taobao.alitrip.car.order.refund

用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
*/
type TaobaoAlitripCarOrderRefundAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripCarOrderRefundAPIResponseModel
}

// 用户投诉达成一致后给用户退款 成功返回结果
type TaobaoAlitripCarOrderRefundAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_car_order_refund_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // 错误消息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 错误码
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}
