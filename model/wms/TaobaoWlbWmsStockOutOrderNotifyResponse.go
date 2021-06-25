package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
出库单通知 APIResponse
taobao.wlb.wms.stock.out.order.notify

出库单通知
*/
type TaobaoWlbWmsStockOutOrderNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWmsStockOutOrderNotifyResponse `json:"taobao_wlb_wms_stock_out_order_notify_response,omitempty"`
}

type TaobaoWlbWmsStockOutOrderNotifyResponse struct {

    // 错误编码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误详细
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 仓储订单编码
    OrderCode   string `json:"order_code,omitempty"`

}
