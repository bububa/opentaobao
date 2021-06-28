package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送订单 APIResponse
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送
*/
type AlibabaEleFengniaoOrderPushAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_order_push_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // msg
    
    Message   string `json:"message,omitempty" xml:"