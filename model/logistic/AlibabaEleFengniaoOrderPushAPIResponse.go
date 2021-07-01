package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送订单 API返回值 
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送
*/
type AlibabaEleFengniaoOrderPushAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoOrderPushAPIResponseModel
}

// 推送订单 成功返回结果
type AlibabaEleFengniaoOrderPushAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_order_push_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // msg
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
