package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推送订单 APIResponse
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送
*/
type AlibabaEleFengniaoOrderPushAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoOrderPushResponse `json:"alibaba_ele_fengniao_order_push_response,omitempty"`
}

type AlibabaEleFengniaoOrderPushResponse struct {

    // msg
    Message   string `json:"message,omitempty"`

}
