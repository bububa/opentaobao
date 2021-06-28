package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询运单事件信息 APIResponse
alibaba.ele.fengniao.shippingorder.event

查询运单事件信息
*/
type AlibabaEleFengniaoShippingorderEventAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaEleFengniaoShippingorderEventResponse `json:"alibaba_ele_fengniao_shippingorder_event_response,omitempty"` 
    AlibabaEleFengniaoShippingorderEventResponse
}

/* model for simplify = false
type AlibabaEleFengniaoShippingorderEventResponse struct {

    // shippingOrderEvents
    
    ShippingOrderEvents  struct {
        ShippingOrderEvent  []ShippingOrderEvent `json:"shipping_order_event,omitempty"`
    } `json:"shipping_order_events,omitempty"`
    

    // 终态时间
    
    FinishAt   int64 `json:"finish_at,omitempty"`
    

    // MERCHANT_CANCEL:商家取消,DELIVERY_TIMEOUT:配送超时，系统标记异常
    
    ShippingRemarkCode   string `json:"shipping_remark_code,omitempty"`
    

    // 骑手预计送达时间
    
    PredictDeliveryAt   int64 `json:"predict_delivery_at,omitempty"`
    

}
*/

type AlibabaEleFengniaoShippingorderEventResponse struct {

    // shippingOrderEvents
    ShippingOrderEvents   []ShippingOrderEvent `json:"shipping_order_events,omitempty"`

    // 终态时间
    FinishAt   int64 `json:"finish_at,omitempty"`

    // MERCHANT_CANCEL:商家取消,DELIVERY_TIMEOUT:配送超时，系统标记异常
    ShippingRemarkCode   string `json:"shipping_remark_code,omitempty"`

    // 骑手预计送达时间
    PredictDeliveryAt   int64 `json:"predict_delivery_at,omitempty"`

}
