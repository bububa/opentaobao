package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售退货通知 APIResponse
taobao.wlb.wms.return.order.notify

销售退货通知
*/
type TaobaoWlbWmsReturnOrderNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsReturnOrderNotifyResponse `json:"taobao_wlb_wms_return_order_notify_response,omitempty"`
}

type TaobaoWlbWmsReturnOrderNotifyResponse struct {

    // 订单创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // LBX929829111
    StoreOrderCode   string `json:"store_order_code,omitempty"`

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

}
