package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货订单通知 APIResponse
taobao.wlb.wms.consign.order.notify

发货订单通知
*/
type TaobaoWlbWmsConsignOrderNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsConsignOrderNotifyResponse `json:"wlb_wms_consign_order_notify_response,omitempty"` 
    TaobaoWlbWmsConsignOrderNotifyResponse
}

/* model for simplify = false
type TaobaoWlbWmsConsignOrderNotifyResponse struct {

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 错误详细
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 订单编码
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 系统自动生成
    
    ConsignOrderList  struct {
        Consignorderlist  []Consignorderlist `json:"consignorderlist,omitempty"`
    } `json:"consign_order_list,omitempty"`
    

}
*/

type TaobaoWlbWmsConsignOrderNotifyResponse struct {

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误详细
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 订单编码
    OrderCode   string `json:"order_code,omitempty"`

    // 系统自动生成
    ConsignOrderList   []Consignorderlist `json:"consign_order_list,omitempty"`

}
