package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运单事件信息 APIResponse
alibaba.ele.fengniao.shippingorder.event

查询运单事件信息
*/
type AlibabaEleFengniaoShippingorderEventAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_shippingorder_event_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // shippingOrderEvents
    
    ShippingOrderEvents   []ShippingOrderEvent `json:"shipping_order_events,omitempty" xml:"