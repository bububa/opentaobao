package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 APIResponse
taobao.wlb.wms.order.cancel.notify

单据取消接口
*/
type TaobaoWlbWmsOrderCancelNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsOrderCancelNotifyResponse `json:"wlb_wms_order_cancel_notify_response,omitempty"` 
    TaobaoWlbWmsOrderCancelNotifyResponse
}

/* model for simplify = false
type TaobaoWlbWmsOrderCancelNotifyResponse struct {

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 错误详细
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

}
*/

type TaobaoWlbWmsOrderCancelNotifyResponse struct {

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误详细
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

}
