package ship

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息回填(出票回调) API返回值 
alitrip.ship.order.notify

此接口为接入商调用飞猪旅行接口回填票号、密码(验证码)等订单信息。接口根据alitripOrderId幂等。若第一次调用失败，后续调用仍然可以回填票号、密码(验证码)成功。第一次调用成功后，后续调用会直接返回第一次的调用结果，不会再产生更新操作。多张票同时出票回填时，保证原子性，只允许全部成功或者全部失败，不能存在部分成功或者失败
*/
type AlitripShipOrderNotifyAPIResponse struct {
    model.CommonResponse
    AlitripShipOrderNotifyAPIResponseModel
}

// 订单信息回填(出票回调) 成功返回结果
type AlitripShipOrderNotifyAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_ship_order_notify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // 错误描述
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
    // 是否成功
    RetSuccess   bool `json:"ret_success,omitempty" xml:"ret_success,omitempty"`
}
